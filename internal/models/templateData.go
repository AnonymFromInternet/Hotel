package models

import "github.com/anonymfrominternet/Hotel/internal/forms"

// TemplateData is a datatype that gives some sort of data into the render function and those data are available in every page
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	CSRFToken string
	Message   string
	Warning   string
	Error     string
	Form      *forms.Form
}
