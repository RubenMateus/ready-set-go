package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type SearchPage struct {
	Query string
	Books []Book
}

type Book struct {
	Title  string
	Author string
	Year   int
}

func main() {
	templates := template.Must(template.ParseFiles("templates/index.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if e := templates.ExecuteTemplate(w, "index.html", nil); e != nil {
			http.Error(w, e.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		books := []Book{
			Book{"Clean Code", "Robert C. Martin", 2008},
			Book{"The Clean Coder", "Robert C. Martin", 2011},
			Book{"An Introduction to programming in GO", "Caleb Doxsey", 2012},
		}
		p := SearchPage{Query: r.FormValue("search"), Books: books}

		if err := templates.ExecuteTemplate(w, "index.html", p); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println(http.ListenAndServe(":504", nil))
}
