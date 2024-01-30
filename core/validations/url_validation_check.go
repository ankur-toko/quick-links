package validations

import (
	"errors"
	"net/url"

	"github.com/ankur-toko/quick-links/core/models"
)

type URLRegexCheck struct {
}

func init() {
	addValidation(CreateURLRegexCheck())
}

func CreateURLRegexCheck() URLRegexCheck {
	return URLRegexCheck{}
}

func (urlChecker URLRegexCheck) Check(r models.QuickLink) error {
	if !urlChecker.isURL(r.URL) {
		return errors.New("incorrect url format")
	}
	return nil
}

func (URLRegexCheck) isURL(url string) bool {
	return IsUrl(url)
}

func IsUrl(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}
