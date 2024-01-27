package validations

import (
	"errors"

	"github.com/ankur-toko/quick-links/core/models"
)

type URLValidator struct{}

func (URLValidator) Check(r models.QuickLink) error {
	if len(r.URL) <= 5 {
		return errors.New("url should be atleast 6 characters long")
	}

	if len(r.Key) > 10 {
		return errors.New("key should not be more than 10 characters long")
	}
	return nil
}
