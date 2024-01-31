package validations

import (
	"testing"

	"github.com/ankur-toko/quick-links/core/models"
)

// Add more tests for this validation
func TestURLRegexCheck_Check(t *testing.T) {
	type args struct {
		r models.QuickLink
	}

	urlChecker := CreateURLRegexCheck()

	tests := []struct {
		name       string
		urlChecker URLRegexCheck
		args       args
		wantErr    bool
	}{
		{
			"URL Validation 1",
			urlChecker,
			args{
				models.QuickLink{Key: "abc", URL: "www.google.com"},
			},
			true,
		},
		{
			"URL Validation 2",
			urlChecker,
			args{
				models.QuickLink{Key: "abc", URL: "google.com"},
			},
			true,
		},
		{
			"URL Validation 3",
			urlChecker,
			args{
				models.QuickLink{"abc", "http://www.google.com"},
			},
			false,
		},
		{
			"URL Validation 4",
			urlChecker,
			args{
				models.QuickLink{Key: "abc", URL: "www.google"},
			},
			true,
		},
		{
			"URL Validation 5",
			urlChecker,
			args{
				models.QuickLink{Key: "abc", URL: "http://www.google"},
			},
			false,
		},
		{
			"URL Validation 6",
			urlChecker,
			args{
				models.QuickLink{Key: "abc", URL: "http://10.41.4.1"},
			},
			false,
		},
		{
			"URL Validation 7",
			urlChecker,
			args{
				models.QuickLink{Key: "abc", URL: "http://10.41.4.1?asd=afada&asdlijasd=afnalkfsnsladmnaksdmasdamsd"},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			urlChecker := URLRegexCheck{}
			if err := urlChecker.Check(tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("URLRegexCheck.Check() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
