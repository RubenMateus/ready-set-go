package services

type Book struct {
	ID     string
	Isbn   string
	Title  string
	Author *Author
}
