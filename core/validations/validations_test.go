package validations

import (
	"testing"
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
