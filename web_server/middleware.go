package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// los middlewares son handlers que reciben handlers para evaluar cierta logica

//CheckAuth ...
func CheckAuth() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			flag := true
			fmt.Println("Checking Authentication")
			if flag {
				f(w, r)
			} else {
				return
			}
		}
	}
}

//Logging ...
func Logging() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			defer func() { //funcion anonima por que la hemos definido y llamado hay mismo sin asignarle un nomre
				log.Println(r.URL.Path, time.Since(start))
			}() // para este tipo de funciones se colocan los parentesis despues de la llave para finalizar
			f(w, r) // con esto el middleware salta al siguiente
		}
	}
}
