package render

import (
	"encoding/gob"
	"github.com/alexedwards/scs/v2"
	"github.com/anonymfrominternet/Hotel/internal/config"
	"github.com/anonymfrominternet/Hotel/internal/models"
	"net/http"
	"os"
	"testing"
	"time"
)

var session *scs.SessionManager
var testAppConfig config.AppConfig

func TestMain(m *testing.M) {
	// Adding custom data types to scs.SessionManager
	gob.Register(models.ReservationPageInputtedData{})
	// Adding custom data types to scs.SessionManager

	// State Management configuration
	session = scs.New()
	session.Lifetime = 3 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false
	// State Management configuration

	// AppConfig and Repository configuration
	testAppConfig.IsInProduction = false
	testAppConfig.Session = session

	appConfig = &testAppConfig

	os.Exit(m.Run())
}

// Adapting custom data type for http.ResponseWriter interface by implementing methods of this interface
type RsWriter struct{}

func (rsWriter *RsWriter) Header() http.Header {
	var header http.Header
	return header
}

func (rsWriter *RsWriter) Write(sliceOfBytes []byte) (int, error) {
	length := len(sliceOfBytes)
	return length, nil
}

func (rsWriter *RsWriter) WriteHeader(statusCode int) {}
