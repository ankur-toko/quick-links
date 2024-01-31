package config

import (
	"fmt"
	"os"
	"testing"
)

func createFile(filename string) {
	contents := `randomkey1: randomVal1
randomkey2: randomVal2
	`
	e := os.WriteFile(filename, []byte(contents), os.ModePerm)
	if e != nil {
		fmt.Print("unable to create file for running tests")
	}
}

func createImproperFile(filename string) {
	contents := `: randomVal1
randomkey1: randomVal1`
	e := os.WriteFile(filename, []byte(contents), os.ModePerm)
	if e != nil {
		fmt.Print("unable to create file for running tests")
	}
}

func deleteFiles(filenames ...string) {
	for _, f := range filenames {
		os.Remove(f)
	}
}

func TestFileConfig_Reload(t *testing.T) {
	testFileName := "correctFile.yaml"
	createFile(testFileName)

	improperFormatFile := "improperFormat.yaml"
	createImproperFile(improperFormatFile)

	defer deleteFiles(testFileName, improperFormatFile)

	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		f       *FileConfig
		args    args
		wantErr bool
	}{
		// Ignore if file is not present
		{
			"FileConfig incorrect",
			&FileConfig{},
			args{
				filename: "afliafjawf",
			},
			false,
		},
		{
			"FileConfig present",
			&FileConfig{},
			args{
				filename: testFileName,
			},
			false,
		},
		{
			"FileConfig improper present",
			&FileConfig{},
			args{
				filename: improperFormatFile,
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &FileConfig{}
			if err := f.ReloadFile(tt.args.filename); (err != nil) != tt.wantErr {
				t.Errorf("FileConfig.Reload() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFileConfig_Get(t *testing.T) {
	testFileName := "correctFile.yaml"
	createFile(testFileName)
	defer deleteFiles(testFileName)

	fileConfig := FileConfig{}
	fileConfig.ReloadFile(testFileName)

	// Set is also tested here
	fileConfig.Set("manualKey1", "manualVal1")

	type args struct {
		key string
	}
	tests := []struct {
		name  string
		fc    *FileConfig
		args  args
		want  string
		want1 bool
	}{
		{
			"Get Test 1",
			&fileConfig,
			args{
				"randomkey1",
			},
			"randomVal1",
			true,
		},
		{
			"Get Test 2",
			&fileConfig,
			args{
				"randomkey123",
			},
			"",
			false,
		},
		{
			"Get Test 3",
			&fileConfig,
			args{
				"manualKey1",
			},
			"manualVal1",
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fc := tt.fc
			got, got1 := fc.Get(tt.args.key)
			if got != tt.want {
				t.Errorf("FileConfig.Get() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("FileConfig.Get() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
