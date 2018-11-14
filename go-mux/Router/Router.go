package Router

import (
	"net/http"

	"github.com/gorilla/mux"
	_logger "github.com/rubenmateus/ready-set-go/go-mux/Logger"
)

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = _logger.Log(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}