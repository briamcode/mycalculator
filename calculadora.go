package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin) //creamos una nueva instancia de escaneo
	scanner.Scan()                        //se realza el escaneo del input
	operacion := scanner.Text()           //creamos una variable llamda operacion que contendra el input del scanner
	fmt.Println(operacion)
	valores := strings.Split(operacion, "+") //usamos Split para fraccionar el string usando + como separador y convirtiendo los datos en 2 arrays
	fmt.Println(valores)
	fmt.Println(valores[0] + valores[1])     //sumamos los arrays [1] Y [2] pero su tipo de dato es string por lo que no los sumara si no que los concatena
	operador1, _ := strconv.Atoi(valores[0]) //aplicamos una transformacion con Atoi para convertir el string a un entero
	operador2, _ := strconv.Atoi(valores[1]) //usamos _ guion bajo para que go nos permita usar la variable de esta forma

	fmt.Println(operador1 + operador2) // finalmente hacemos la operacion aritmetica y la imprimimos
}
