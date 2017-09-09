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

