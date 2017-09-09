package viewModels

type Books struct {
	Title string
	Img string
	Categories[] Category
	Description string
}

type Category struct {
	Name string
}

func NewBooks() Books {
	cat1 := Category{Name:"Biography"}
	return Books{
		Title: "Elon Musk",
		Img: "https://www.manufacturing.net/sites/manufacturing.net/files/elon%20musk%20book%20jacket%202.jpg",
		Description: "Elon Musk: Tesla, SpaceX, and the Quest for a Fantastic Future is Ashlee Vance's biography of Elon Musk, published in 2015...",
		Categories: []Category{cat1},
	}
}

