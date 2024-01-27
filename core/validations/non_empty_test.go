package validations

import (
	"testing"

	"github.com/ankur-toko/quick-links/core/models"
)

func TestNonEmptyCheck_Check(t *testing.T) {
	type args struct {
		r models.QuickLink
	}
	tests := []struct {
		name    string
		n       NonEmptyCheck
		args    args
		wantErr bool
	}{
		{
			"NonEmpty Check 1",
			NonEmptyCheck{},
			args{
				models.QuickLink{Key: "", URL: "adasd"},
			},
			true,
		},
		{
			"NonEmpty Check 2",
			NonEmptyCheck{},
			args{
				models.QuickLink{Key: "adaf", URL: ""},
			},
			true,
		},
		{
			"NonEmpty Check 3",
			NonEmptyCheck{},
			args{
				models.QuickLink{Key: "", URL: ""},
			},
			true,
		},
		{
			"NonEmpty Check 4",
			NonEmptyCheck{},
			args{
				models.QuickLink{Key: "asdad", URL: "afkdsjhsdkla"},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := NonEmptyCheck{}
			if err := n.Check(tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("NonEmptyCheck.Check() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
