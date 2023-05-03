package model

type Cookie struct {
	Domain         string
	ExpirationDate int
	HttpOnly       bool
	Name           interface{}
	Path           interface{}
	Secure         bool
	Value          string
}
