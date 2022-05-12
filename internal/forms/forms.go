package forms

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type Form struct {
	url.Values
	Errors errors
}

func (form *Form) Valid() bool {
	return len(form.Errors) == 0
}

// New initializes new object of struct Form
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// Required checks if a form field (input) is required
func (form *Form) Required(fields ...string) {
	for _, field := range fields {
		value := form.Get(field)
		if strings.TrimSpace(value) == "" {
			form.Errors.Add(field, "This field cannot be empty")
		}
	}
}

// MinLength checks if inputted data has required length
func (form *Form) MinLength(field string, length int, request *http.Request) bool {
	value := request.Form.Get(field)
	if len(value) < length {
		form.Errors.Add(field, fmt.Sprintf("Entered length must have minimum %d characters", length))
		return false
	}
	return true
}

// HasFieldValue checks if form field was send in POST method and is not empty
// should be used in the future
func (form *Form) HasFieldValue(field string, request *http.Request) bool {
	inputtedValueFromFormField := request.Form.Get(field)
	if inputtedValueFromFormField == "" {
		return false
	}
	return true
}
