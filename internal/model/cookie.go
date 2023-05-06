package model

type Cookie struct {
	Domain         string
	ExpirationDate float64
	HttpOnly       interface{}
	Name           string
	Path           string
	Secure         interface{}
	Value          string
	Key            string
	Expires        float64
}
