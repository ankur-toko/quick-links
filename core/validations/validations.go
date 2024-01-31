package validations

import (
	"errors"
	"strings"

	"github.com/ankur-toko/quick-links/core/models"
)

var validations []ValidationCheck

type ValidationCheck interface {
	// Returns nil if no problem
	Check(models.QuickLink) error
}

type Validations struct {
	validators []ValidationCheck
}

func (v *Validations) add(checks ValidationCheck) {
	v.validators = append(v.validators, checks)
}

func (v Validations) Check(r models.QuickLink) error {
	errs := []string{}
	for _, v := range v.validators {
		e := v.Check(r)
		if e != nil {
			errs = append(errs, e.Error())
		}
	}
	if len(errs) > 0 {
		return errors.New(strings.Join(errs, ","))
	}
	return nil
}

func DefaultValidatorObj() Validations {
	return Validations{validators: validations}
}

func addValidation(v ValidationCheck) {
	validations = append(validations, v)
}
