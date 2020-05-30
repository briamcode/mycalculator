package main

import "fmt"

func main() {
	x := 25
	fmt.Println(x)
	cambiarValor(x)
	fmt.Println(x)
	y := &x
	fmt.Println(y)
	fmt.Println(*y) // se trae el valor de esa direccion especifica de memoria

}

func cambiarValor(a int) {
	fmt.Println(&a) // se usa ampersan para mostrar la direccion en memoria
	a = 36
	fmt.Println(a)
}
