package workers

import (
	"encoding/csv"
	"fmt"
	"github.com/MorZLE/GoParseTSV/internal/model"
	"os"
	"reflect"
	"testing"
)

func TestParser_Parse(t *testing.T) {

	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		want    []model.Guid
		wantSl  []string
		wantErr bool
	}{
		{
			name: "positiveTest1",
			args: args{
				filename: "test2.tsv",
			},
			want: []model.Guid{
				{
					Number:       "n",
					MQTT:         "mqtt",
					InventoryID:  "invid",
					UnitGUID:     "unit_guid",
					MessageText:  "text",
					MessageID:    "msg_id",
					Context:      "context",
					MessageClass: "class",
					Level:        "level",
					Area:         "area",
					Address:      "addr",
					Block:        "block",
					Type:         "type",
					Bit:          "bit",
					InvertBit:    "invert_bit",
				},
				{
					Number:       "1",
					InventoryID:  "G-044322",
					UnitGUID:     "01749246-95f6-57db-b7c3-2ae0e8be671f",
					MessageText:  "Разморозка",
					MessageID:    "cold7_Defrost_status",
					MessageClass: "waiting",
					Level:        "100",
					Area:         "LOCAL",
					Address:      "cold7_status.Defrost_status",
				},
			},
			wantSl: []string{
				"01749246-95f6-57db-b7c3-2ae0e8be671f",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testdir := "testdir"
			createTempDir(t, testdir)
			defer removeTempDir(testdir)

			createFile(testdir, tt.args.filename, tt.want)
			s := &Parser{
				dirIN: testdir,
			}
			got, got1, err := s.Parse(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.wantSl) {
				t.Errorf("Parse() got1 = %v, want %v", got1, tt.wantSl)
			}
		})
	}
}

// createTempDir создает временную директорию
func createTempDir(t *testing.T, dir string) {
	err := os.Mkdir(dir, os.ModeDir)
	if err != nil {
		t.Fatalf("Failed to create temporary directory: %v", err)
	}
	err = os.Chmod(dir, 0700)
	if err != nil {
		t.Fatalf("Failed to create Chmod directory: %v", err)
	}
}

// removeTempDir удаляет временную директорию
func removeTempDir(dir string) {
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

func createFile(dir string, name string, model []model.Guid) {
	// Specify the file path
	filePath := dir + "/" + name

	// Create the file
	file, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Create a CSV writer
	writer := csv.NewWriter(file)
	writer.Comma = '\t'

	// Write each Person struct as a row in the TSV file
	for _, m := range model {
		err = writer.Write([]string{
			m.Number,
			m.MQTT,
			m.InventoryID,
			m.UnitGUID,
			m.MessageID,
			m.MessageText,
			m.Context,
			m.MessageClass,
			m.Level,
			m.Area,
			m.Address,
			m.Block,
			m.Type,
			m.Bit,
			m.InvertBit,
		})
		if err != nil {
			panic(err)
		}
	}

	writer.Flush()
	if err := writer.Error(); err != nil {
		panic(err)
	}
	fmt.Println("TSV file created and written successfully.")
}
