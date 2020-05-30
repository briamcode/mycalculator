package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	inicio := time.Now()
	canal := make(chan string)

	servidores := []string{
		"http://www.platzi.com",
		"http://www.google.com",
		"http://www.facebook.com",
		"http://www.instagram.com",
		"http://www.dian.gov.co",
	}

	i := 0

	for {
		if i > 10 {
			break
		}
		for _, servidor := range servidores {
			go revisarServidor(servidor, canal)

		}
		time.Sleep(1 * time.Second)
		fmt.Println(<-canal)
		i++
	}

	tiempoPaso := time.Since(inicio)
	fmt.Printf("Tiempo de ejecucion %s", tiempoPaso)
}

func revisarServidor(servidor string, canal chan string) {
	_, err := http.Get(servidor)
	if err != nil {

		canal <- servidor + "no se encuentra disponible"
	} else {

		canal <- servidor + " esta funcionando"
	}
}
