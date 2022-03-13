package models

import "github.com/anonymfrominternet/Hotel/internal/forms"

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