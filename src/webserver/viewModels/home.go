package viewModels

type Base struct {
	Title string
}

func NewBase() Base {
	return Base{
		Title: "WebServer",
	}
}

