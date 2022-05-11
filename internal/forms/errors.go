package forms

type errors map[string][]string

// Add adds an error message for a given form field
func (errors errors) Add(field string, message string) {
	errors[field] = append(errors[field], message)
}

// Get returns the first error message
func (errors errors) Get(field string) string {
	errorString := errors[field]
	if len(errorString) == 0 {
		return ""
	}
	return errorString[0]
}
