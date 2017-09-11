package viewModels

import "webserver/model"

type BookVm struct {
	Title string
	Books []model.Book
}
func NewBooks(books []model.Book) BookVm {
	return BookVm{
		Title: "Awesome Books",
		Books: books,
	}
}

type Category struct {
	Name string
}

type Book struct {
	Id int `json: "id"`
	Title string `json: "title"`
	Img string `json: "img"`
	Categories[] Category `json: "categories"`
	Description string `json: "description"`
	Published string `json: "published"`
}