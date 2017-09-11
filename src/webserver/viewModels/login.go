package viewModels

type Login struct {
	Email string
	Password string
}


func NewLogin() Login {
	return Login{}
}