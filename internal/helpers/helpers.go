package helpers

import (
	"fmt"
	"github.com/anonymfrominternet/Hotel/internal/config"
	"net/http"
	"runtime/debug"
)

var appConfig *config.AppConfig

// GetAppConfigToTheHelpersPackage brings appConfig from main package into this package
func GetAppConfigToTheHelpersPackage(appConfigAsParam *config.AppConfig) {
	appConfig = appConfigAsParam
}

func ClientError(writer http.ResponseWriter, status int) {
	appConfig.InfoLog.Println("Client error with status", status)

	// Message for user
	http.Error(writer, http.StatusText(status), status)

}

func ServerError(writer http.ResponseWriter, err error) {
	// trace holds detailed info with stack about errors
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	appConfig.ErrorLog.Println(trace)

	// Message for user
	http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func IsAuthenticated(request *http.Request) bool {
	exists := appConfig.Session.Exists(request.Context(), "user_id")
	return exists
}
