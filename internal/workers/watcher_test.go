package workers

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestWatcher_Scan(t *testing.T) {

	tests := []struct {
		name string
	}{
		{name: "test1"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			createTempDir(t)
			defer removeTempDir()

			createTempFile(t, "file1.tsv")
			createTempFile(t, "file2.tsv")

			w := &Watcher{
				dirIN:      "testdir",
				tickerTime: 1,
				fileCheck:  make(map[string]bool),
			}

			out := make(chan string)
			defer close(out)
			//cancel := make(chan struct{})
			//defer close(cancel)

			go w.Scan(out)

			// Both files should be sent to the output channel
			select {
			case file := <-out:
				if file != "file1.tsv" && file != "file2.tsv" {
					t.Errorf("Expected file1.tsv or file2.tsv to be sent, got %s", file)
				}
			case <-time.After(10 * time.Second):
				t.Errorf("Expected file1.tsv or file2.tsv to be sent")
			}
			select {
			case file := <-out:
				if file != "file1.tsv" && file != "file2.tsv" {
					t.Errorf("Expected file1.tsv or file2.tsv to be sent, got %s", file)
				}
			case <-time.After(10 * time.Second):
				t.Errorf("Expected file1.tsv or file2.tsv to be sent")
			}
		})
	}
	removeTempDir()
}

// Helper function to create a temporary directory
func createTempDir(t *testing.T) {
	err := os.Mkdir("testdir", os.ModeDir)
	if err != nil {
		t.Fatalf("Failed to create temporary directory: %v", err)
	}
	err = os.Chmod("testdir", 0700)
	if err != nil {
		t.Fatalf("Failed to create Chmod directory: %v", err)
	}
}

// Helper function to remove a temporary directory

func removeTempDir() {
	dir := "testdir"
	projectDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	s := fmt.Sprintf("%s\\%s\\", projectDir, dir)
	err = os.RemoveAll(s)
	if err != nil {
		panic(err)
	}
}

// Helper function to create a temporary file
func createTempFile(t *testing.T, name string) {
	dir := "testdir"
	filePath := filepath.Join(dir, name)
	f, err := os.Create(filePath)
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	f.Close()
}
