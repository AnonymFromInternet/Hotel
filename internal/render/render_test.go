package render

import (
	"github.com/anonymfrominternet/Hotel/internal/models"
	"net/http"
	"testing"
)

func TestAddDefaultData(t *testing.T) {

	var templateData models.TemplateData
	// Creating new request with a context
	request, err := getSession()
	if err != nil {
		t.Error(err)
	}
	session.Put(request.Context(), "Message", "message")
	// Get data from AppConfig for pages by the rendering
	result := AddDefaultData(&templateData, request)

	if result.Message != "message" {
		t.Error("Value of Message is not equal message")
	}
}

func getSession() (*http.Request, error) {
	request, err := http.NewRequest("GET", "/url", nil)
	if err != nil {
		return nil, err
	}

	context := request.Context()
	context, _ = session.Load(context, request.Header.Get("X-Session"))
	request = request.WithContext(context)

	return request, nil
}

func TestTemplate(t *testing.T) {
	templateCache, err := CreateTemplateCache()
	if err != nil {
		t.Error(err)
	}
	appConfig.TemplateCache = templateCache

	var rw RsWriter
	request, err := getSession()
	if err != nil {
		t.Error(err)
	}
	err = Template(&rw, request, "main.page.tmpl", &models.TemplateData{})
	if err != nil {
		t.Error(err)
	}
	err = Template(&rw, request, "not-existing.page.tmpl", &models.TemplateData{})
	if err == nil {
		t.Error("this page cannot exist: ", err)
	}
}

func TestNewTemplates(t *testing.T) {
	NewTemplates(appConfig)
}

func TestCreateTemplateCache(t *testing.T) {
	_, err := CreateTemplateCache()
	if err != nil {
		t.Error(err)
	}
}
