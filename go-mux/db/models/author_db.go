package db

import "github.com/jinzhu/gorm"

type Author struct {
	gorm.Model
	FirstName string
	LastName  string
}
