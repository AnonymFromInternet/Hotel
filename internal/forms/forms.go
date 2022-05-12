package forms

import (
	"net/http"
	"net/url"
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

// HasFieldValue checks if form field was send in POST method and is not empty
func (form *Form) HasFieldValue(field string, request *http.Request) bool {
	inputtedValueFromFormField := request.Form.Get(field)
	if inputtedValueFromFormField == "" {
		form.Errors.Add(field, "This field cannot be empty")
		return false
	}
	return true
}
