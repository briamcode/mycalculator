package main

import (
	"encoding/json"
	"net/http"
)

//Middleware ...
type Middleware func(http.HandlerFunc) http.HandlerFunc

//MetaData ...
type MetaData interface {
}

//User ...
type User struct {
	Name  string
	Email string
	Phone string
}

//ToJson ...
func (u *User) ToJson() ([]byte, error) {
	return json.Marshal(u)
}
