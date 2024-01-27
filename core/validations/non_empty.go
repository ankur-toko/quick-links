package validations

import (
	"errors"

	"github.com/ankur-toko/quick-links/core/models"
)

type NonEmptyCheck struct{}

func (NonEmptyCheck) Check(r models.QuickLink) error {
	if r.Key == "" || r.URL == "" {
		return errors.New("key or url cannot be empty strings")
	}
	return nil
}
