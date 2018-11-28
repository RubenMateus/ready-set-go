package db

import "github.com/jinzhu/gorm"

type Book struct {
	gorm.Model
	ID       string
	Isbn     string
	Title    string
	AuthorID uint
}
