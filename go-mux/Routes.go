package main

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{"GetBooks", "GET", "/api/books", GetBooks},
	Route{"GetBook", "GET", "/api/books/{id}", GetBook},
	Route{"CreateBook", "POST", "/api/books", CreateBook},
	Route{"UpdateBook", "PUT", "/api/books/{id}", UpdateBook},
	Route{"DeleteBook", "DELETE", "/api/books/{id}", DeleteBook},
}
