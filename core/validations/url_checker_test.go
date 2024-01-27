package validations

import (
	"testing"

	"github.com/ankur-toko/quick-links/core/models"
)

func TestURLValidator_Check(t *testing.T) {
	type args struct {
		r models.QuickLink
	}
	tests := []struct {
		name    string
		u       URLValidator
		args    args
		wantErr bool
	}{
		{
			"URL Checker 1",
			URLValidator{},
			args{
				models.QuickLink{Key: "adaf", URL: ""},
			},
			true,
		},
		{
			"URL Checker 2",
			URLValidator{},
			args{
				models.QuickLink{Key: "adaf", URL: "asda"},
			},
			true,
		},
		{
			"URL Checker",
			URLValidator{},
			args{
				models.QuickLink{Key: "adaf", URL: "laada;sldka;sdlka;"},
			},
			false,
		},
		{
			"URL Checker",
			URLValidator{},
			args{
				models.QuickLink{Key: "addaasdasdasd", URL: "laadasldksdlka"},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := URLValidator{}
			if err := u.Check(tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("URLValidator.Check() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
