package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type calc struct{}

func (calc) operate(val1, val2, operator string) float64 {
	switch operator {
	case "+":
		return parse(val1) + parse(val2)
	case "-":
		return parse(val1) - parse(val2)
	case "*":
		return parse(val1) * parse(val2)
	case "/":
		return parse(val1) / parse(val2)
	default:
		return 0
	}
}

//Lee un valor ingresado por el usuario.
func readInput() string {
	scanner := bufio.NewScanner(os.Stdin) //Creación Scanner para ingreso de datos.
	scanner.Scan()                        //Inicio del scanner
	return scanner.Text()                 //Retornar datos ingresados
}

//Convierte un valor string a Float.
func parse(val string) float64 {
	s, err := strconv.ParseFloat(val, 64)
	if err == nil {
		return s
	}

	fmt.Println(err)
	return 0
}

func main() {
	fmt.Println("=================CALCULADORA==================")
	fmt.Println("Ingrese primer valor:")
	val1 := readInput()
	fmt.Println("Ingrese operador aritmetico (+ - * /)")
	operator := readInput()
	fmt.Println("Ingrese segundo valor:")
	val2 := readInput()
	fmt.Println(val1 + operator + val2)
	miCalc := calc{}
	fmt.Printf("El resultado es: %v\n", miCalc.operate(val1, val2, operator))
	readInput()
}
