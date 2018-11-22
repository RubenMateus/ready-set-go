package router

import h "github.com/rubenmateus/ready-set-go/go-mux/handlers"

var routes = Routes{
	Route{"GetBooks", "GET", "/api/books", h.GetBooks},
	Route{"GetBook", "GET", "/api/books/{id}", h.GetBook},
	Route{"CreateBook", "POST", "/api/books", h.CreateBook},
	Route{"UpdateBook", "PUT", "/api/books/{id}", h.UpdateBook},
	Route{"DeleteBook", "DELETE", "/api/books/{id}", h.DeleteBook},
}
