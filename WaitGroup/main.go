package main

import (
	"fmt"
	"runtime"
	"sync"
)

//Importa grupo para manejar GORoutines de forma concurrente
var wg sync.WaitGroup

func main() {
	fmt.Println("OS\t", runtime.GOOS)
	fmt.Println("ARCH\t", runtime.GOARCH)
	fmt.Println("CPUs\t", runtime.NumCPU())
	fmt.Println("GoRoutines\t", runtime.NumGoroutine())

	//Se le indica el numero de GORoutines que se vana a utilizar, en este caso solo la de foo.
	wg.Add(1)
	go foo()
	bar()

	fmt.Println("GoRoutines\t", runtime.NumGoroutine())
	//Espera a que finalicen todas las GORoutines.
	wg.Wait()
}

func foo() {
	for i := 0; i < 6; i++ {
		fmt.Println("foo:", i)
	}

	//Se indica que esta GORoutine ya termino.
	wg.Done()
}

func bar() {
	for i := 0; i < 6; i++ {
		fmt.Println("bar:", i)
	}
}
