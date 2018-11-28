package db

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	dbModels "github.com/rubenmateus/ready-set-go/go-mux/db/models"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "123"
	dbname   = "mux-books"
)

func GetBooks() []dbModels.Book {
	dbURL := os.Getenv("DATABASE_URL")

	if dbURL == "" {
		dbURL = fmt.Sprintf(
			"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			host, port, user, password, dbname)
	}

	db, err := gorm.Open("postgres", dbURL)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.AutoMigrate(&dbModels.Book{}, &dbModels.Author{})

	var books []dbModels.Book
	db.Find(&books)

	return books
}
