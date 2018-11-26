package router

import (
	"github.com/gorilla/mux"
	logger "github.com/rubenmateus/ready-set-go/go-mux/utils/logger"
)

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {

		handler := logger.Log(route.HandlerFunc, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}
