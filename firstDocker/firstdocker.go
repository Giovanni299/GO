package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("=================HOLA MUNDO!!!==================")
	fmt.Println("Ingrese su nombre:")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	name := scanner.Text()

	fmt.Println("Hola " + name)
}

/*
Para ejecutarlo:
go run firstdocker.go
1. docker build -t firstdocker .   //Contruir la imagen
2. docker run -it firstdocker   //Ejecutar la aplicación
*/
