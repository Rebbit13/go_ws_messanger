package service

type Authorization interface {
	SignUp(username string, password string) error
	SignIn(username string, password string) error
}
