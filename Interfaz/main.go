package main

import (
	"fmt"
	"math"
)

type circulo struct {
	radio float64
}

type cuadrado struct {
	lado float64
}

//Definir interfaz, para el metodo area
type forma interface {
	area() float64
}

func main() {
	cir := circulo{
		radio: 12.345,
	}

	cuad := cuadrado{
		lado: 15,
	}

	//Area circulo
	info(cir)

	//Area cuadrado
	info(cuad)
}

func (c circulo) area() float64 {
	return math.Pi * c.radio * c.radio
}

func (c cuadrado) area() float64 {
	return c.lado * c.lado
}

//Funcion que recibe una interfaz e imprime el area de la forma (circulo - cuadrado)
func info(f forma) {
	fmt.Println(f.area())
}
