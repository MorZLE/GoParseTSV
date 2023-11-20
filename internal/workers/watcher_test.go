package workers

import (
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
	tesdir := "testdir1"
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			createTempDir(t, tesdir)

			createTempFile(t, tesdir, "file1.tsv")
			createTempFile(t, tesdir, "file2.tsv")

			w := &Watcher{
				dirIN:      tesdir,
				tickerTime: 1,
				fileCheck:  make(map[string]bool),
			}

			out := make(chan string)
			defer close(out)
			go w.Scan(out)

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
	removeTempDir(tesdir)
}

//createTempDir создает временную директорию
//func createTempDir(t *testing.T) {
//	err := os.Mkdir("testdir", os.ModeDir)
//	if err != nil {
//		t.Fatalf("Failed to create temporary directory: %v", err)
//	}
//	err = os.Chmod("testdir", 0700)
//	if err != nil {
//		t.Fatalf("Failed to create Chmod directory: %v", err)
//	}
//}
//
////removeTempDir удаляет временную директорию
//func removeTempDir() {
//	dir := "testdir"
//	projectDir, err := os.Getwd()
//	if err != nil {
//		panic(err)
//	}
//	s := fmt.Sprintf("%s\\%s\\", projectDir, dir)
//	err = os.RemoveAll(s)
//	if err != nil {
//		panic(err)
//	}
//}

// createTempFile создает временный файл
func createTempFile(t *testing.T, dir, name string) {
	filePath := filepath.Join(dir, name)
	f, err := os.Create(filePath)
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	f.Close()
}
