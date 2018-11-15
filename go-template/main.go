package main

import (
	"encoding/xml"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/url"
)

type SearchPage struct {
	Query string
	Books []Book
}

type Book struct {
	Title  string `xml:"title,attr"`
	Author string `xml:"author,attr"`
	Year   int    `xml:"hyr,attr"`
	ID     string `xml:"owi,attr"`
}

func main() {
	templates := template.Must(template.ParseFiles("templates/index.html"))

	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		results, e := search(r.FormValue("search"))
		if e != nil {
			http.Error(w, e.Error(), http.StatusInternalServerError)
		}

		p := SearchPage{Query: r.FormValue("search"), Books: results}
		if err := templates.ExecuteTemplate(w, "index.html", p); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := templates.ExecuteTemplate(w, "index.html", nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println(http.ListenAndServe(":4000", nil))
}

func search(query string) ([]Book, error) {
	var response struct {
		Books []Book `xml:"works>work"`
	}

	body, err := fetch("title=" + url.QueryEscape(query))

	if err != nil {
		return []Book{}, err
	}

	err = xml.Unmarshal(body, &response)
	return response.Books, err
}

func fetch(q string) ([]byte, error) {
	var resp *http.Response
	var err error
	url := "http://classify.oclc.org/classify2/Classify?summary=true&" + q

	if resp, err = http.Get(url); err != nil {
		return []byte{}, err
	}

	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
