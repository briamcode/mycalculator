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

	for _, servidor := range servidores {
		go revisarServidor(servidor, canal)

	}

	for i := 0; i < len(servidores); i++ {
		fmt.Println(<-canal)
	}

	tiempoPaso := time.Since(inicio)
	fmt.Printf("Tiempo de ejecucion %s", tiempoPaso)
}

func revisarServidor(servidor string, canal chan string) {
	_, err := http.Get(servidor)
	if err != nil {
		fmt.Println(servidor, "no esta disponible =(")
		canal <- servidor + "no se encuentra disponible"
	} else {
		fmt.Println(servidor, "esta funcionando normalmente =(")
		canal <- servidor + " esta funcionando"
	}
}
