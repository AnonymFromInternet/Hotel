package forms

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"net/http"
	"net/url"
	"strings"
)

type Form struct {
	url.Values
	Errors errors
}

func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

// NewForm creates new empty form
func NewForm(data url.Values) *Form {
	return &Form{
		// This property holds inputted data from a form
		Values: data,
		// This property holds inputted data from a form

		Errors: map[string][]string{},
	}
}

func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field cannot be empty")
		}
	}

}
func (f *Form) MinLength(field string, minLength int, request *http.Request) bool {
	x := request.Form.Get(field)
	if len(x) < minLength {
		f.Errors.Add(field, fmt.Sprintf("Value is too small"))
		return false
	}
	return true
}

func (f *Form) Has(field string, request *http.Request) bool {
	x := request.Form.Get(field)
	if x == "" {

		return false
	}
	return true
}
func (f *Form) IsEmail(field string) {
	if !govalidator.IsEmail(f.Get(field)) {
		f.Errors.Add(field, "This is not Email")
	}
}
