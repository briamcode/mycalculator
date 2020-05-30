package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	inicio := time.Now()

	servidores := []string{
		"http://www.platzi.com",
		"http://www.google.com",
		"http://www.facebook.com",
		"http://www.instagram.com",
	}

	for _, servidor := range servidores {
		revisarServidor(servidor)

	}

	tiempoPaso := time.Since(inicio)
	fmt.Printf("Tiempo de ejecucion %s", tiempoPaso)
}

func revisarServidor(servidor string) {
	_, err := http.Get(servidor)
	if err != nil {
		fmt.Println(servidor, "no esta disponible =(")
	} else {
		fmt.Println(servidor, "esta funcionando normalmente =(")
	}
}
