package forms

type errors map[string][]string

// Add adds errors to errors, when form is empty
func (e errors) Add(field, message string) {
	e[field] = append(e[field], message)
}

// Get returns existing error, when it exists (was added by method Add())
func (e errors) Get(field string) string {
	es := e[field]
	if len(es) == 0 {
		return ""
	}
	return es[0]
}
