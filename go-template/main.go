package main

import (
	"encoding/xml"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type SearchPage struct {
	Query   string
	Results []Result
}

type Result struct {
	Title  string `xml:"title,attr"`
	Author string `xml:"author,attr"`
	Year   int    `xml:"hyr,attr"`
	ID     string `xml:"owi,attr"`
}

type Book struct {
	gorm.Model
	Title          string
	Author         string
	OWI            string
	Classification string
}

type BookResponse struct {
	BookData struct {
		Title  string `xml:"title,attr"`
		Author string `xml:"author,attr"`
		ID     string `xml:"owi,attr"`
	} `xml:"work"`
	Classification struct {
		MostPopular string `xml:"sfa,attr"`
	} `xml:"recommendations>ddc>mostPopular"`
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "123"
	dbname   = "go-books"
)

func main() {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.AutoMigrate(&Book{})

	libraryTemplates := template.Must(
		template.ParseFiles("templates/layout.html", "templates/library.html"))
	searchTemplates := template.Must(
		template.ParseFiles("templates/layout.html", "templates/search.html"))

	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "book.ico")
	})

	http.HandleFunc("/removebook", func(w http.ResponseWriter, r *http.Request) {
		var book Book
		db.Find(&book, r.FormValue("bookId"))
		db.Delete(book)

		http.Redirect(w, r, "/", 302)
	})

	http.HandleFunc("/addbook", func(w http.ResponseWriter, r *http.Request) {
		res, e := find(r.FormValue("bookId"))
		if e != nil {
			http.Error(w, e.Error(), http.StatusInternalServerError)
		}
		db.Create(&Book{
			Title:          res.BookData.Title,
			Author:         res.BookData.Author,
			OWI:            res.BookData.ID,
			Classification: res.Classification.MostPopular,
		})

		http.Redirect(w, r, "/", 302)
	})

	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		results, e := search(r.FormValue("search"))
		if e != nil {
			http.Error(w, e.Error(), http.StatusInternalServerError)
		}

		p := SearchPage{Query: r.FormValue("search"), Results: results}
		if err := searchTemplates.ExecuteTemplate(w, "layout", p); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		var p struct{ Books []Book }

		order := r.FormValue("sort")
		if order != "title" && order != "author" && order != "classification" {
			order = "title"
		}

		if filterInt, err := strconv.Atoi(r.FormValue("filter")); err == nil {
			db.Order(order).Where("classification BETWEEN ? AND ?",
				r.FormValue("filter"), strconv.Itoa(filterInt+100)).Find(&p.Books)
		} else {
			db.Order(order).Find(&p.Books)
		}

		if err := libraryTemplates.ExecuteTemplate(w, "layout", p); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println(http.ListenAndServe(":4000", nil))
}

func search(query string) ([]Result, error) {
	var response struct {
		Results []Result `xml:"works>work"`
	}

	body, err := fetch("title=" + url.QueryEscape(query))

	if err != nil {
		return []Result{}, err
	}

	err = xml.Unmarshal(body, &response)
	return response.Results, err
}

func find(id string) (BookResponse, error) {
	var response BookResponse
	body, err := fetch("owi=" + url.QueryEscape(id))

	if err != nil {
		return BookResponse{}, err
	}

	err = xml.Unmarshal(body, &response)
	return response, err
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
