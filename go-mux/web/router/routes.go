package router

import "github.com/rubenmateus/ready-set-go/go-mux/web/controllers"

var routes = Routes{
	Route{"GetBooks", "GET", "/api/books", controllers.GetBooks},
	Route{"GetBook", "GET", "/api/books/{id}", controllers.GetBook},
	Route{"CreateBook", "POST", "/api/books", controllers.CreateBook},
	Route{"UpdateBook", "PUT", "/api/books/{id}", controllers.UpdateBook},
	Route{"DeleteBook", "DELETE", "/api/books/{id}", controllers.DeleteBook},
}
