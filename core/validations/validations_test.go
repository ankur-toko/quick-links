package validations

import (
	"errors"
	"strings"
	"testing"

	"github.com/ankur-toko/quick-links/core/models"
	"github.com/gookit/goutil/testutil/assert"
)

func TestValidations_add(t *testing.T) {
	type fields struct {
		validators []ValidationCheck
	}
	type args struct {
		checks ValidationCheck
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			"Validations Len Check",
			fields{[]ValidationCheck{}},
			args{URLValidator{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &Validations{
				validators: tt.fields.validators,
			}
			v.add(tt.args.checks)

			if len(v.validators) != 1 {
				t.FailNow()
			}
		})
	}
}

type AlwaysErrorCheck struct{}

func (t *AlwaysErrorCheck) Check(models.QuickLink) error {
	return errors.New("some random error")
}

type AlwaysNoErrorCheck struct{}

func (t *AlwaysNoErrorCheck) Check(models.QuickLink) error {
	return nil
}

func TestValidations_Check(t *testing.T) {

	validationObject := Validations{}

	validationObject.validators = []ValidationCheck{&AlwaysErrorCheck{}}
	err := validationObject.Check(models.QuickLink{Key: "key", URL: "value"})
	assert.True(t, err != nil)

	validationObject.validators = []ValidationCheck{&AlwaysNoErrorCheck{}}
	err = validationObject.Check(models.QuickLink{Key: "key", URL: "value"})
	assert.True(t, err == nil)

	validationObject.validators = []ValidationCheck{&AlwaysErrorCheck{}, &AlwaysErrorCheck{}}
	err = validationObject.Check(models.QuickLink{Key: "key", URL: "value"})
	assert.True(t, len(strings.Split(err.Error(), ",")) >= 2)

}

func TestDefaultValidatorObj(t *testing.T) {
	got := DefaultValidatorObj()
	assert.True(t, len(got.validators) > 0)
}
