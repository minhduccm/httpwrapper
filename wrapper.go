package httpwrapper

import (
	"log"
	"net/http"
)

type customHandlerFunc func(http.ResponseWriter, *http.Request) (int, error)

type appHandler struct {
	mainHandler customHandlerFunc
}

func (ah appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	status, err := ah.mainHandler(w, r)
	if err != nil {
		log.Printf("HTTP %d: %q", status, err)
		switch status {
		case http.StatusNotFound:
			http.NotFound(w, r)
		case http.StatusInternalServerError:
			http.Error(w, http.StatusText(status), status)
		default:
			http.Error(w, http.StatusText(status), status)
		}
	}
}

type middlewares []func(http.Handler) http.Handler

func Perform(middlewares ...func(http.Handler) http.Handler) middlewares {
	return middlewares
}

func (middlewares middlewares) LeadTo(handler customHandlerFunc) http.Handler {
	ah := appHandler{handler}

	var temp http.Handler = ah
	for i := len(middlewares) - 1; i >= 0; i-- {
		temp = middlewares[i](temp)
	}
	return temp
}
