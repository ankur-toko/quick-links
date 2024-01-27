package config

import (
	"testing"
)

const test_config_file_path = "../testdata/config/testconfig.yml"

func TestInitializeConfiguration(t *testing.T) {
	type args struct {
		filepath string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"Config File Absent Test",
			args{
				"absent_file_path.json",
			},
			true,
		},
		{
			"Config File Present Test",
			args{
				test_config_file_path,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := InitializeConfiguration(tt.args.filepath); (err != nil) != tt.wantErr {
				t.Errorf("InitializeConfiguration() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGet(t *testing.T) {
	InitializeConfiguration(test_config_file_path)
	type args struct {
		key string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 bool
	}{
		{
			"Get Value present Test",
			args{
				"app.name",
			},
			"test-app-name",
			true,
		},
		{
			"Get Value absent Test",
			args{
				"app.absent_key",
			},
			"",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := Get(tt.args.key)
			if got != tt.want {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Get() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
