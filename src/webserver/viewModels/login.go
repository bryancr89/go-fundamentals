package viewModels

type Login struct {
	email string
	password string
}


func NewLogin() Login {
	return Login{}
}