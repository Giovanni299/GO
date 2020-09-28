package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("OS\t", runtime.GOOS)
	fmt.Println("ARCH\t", runtime.GOARCH)
	fmt.Println("CPUs\t", runtime.NumCPU())
	fmt.Println("GoRoutines\t", runtime.NumGoroutine())

	var wg sync.WaitGroup
	var mutex sync.Mutex
	goRoutines := 10
	totalCount := 0
	wg.Add(goRoutines)

	for i := 0; i < goRoutines; i++ {
		go func() {
			mutex.Lock()
			value := totalCount
			runtime.Gosched() //Yield
			value++
			totalCount = value
			mutex.Unlock()
			wg.Done()
		}()
		fmt.Println("Numero de GoRoutines\t", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("Contador\t", totalCount)
}

//Para ejecutar y validar si hay race-condition
//go run -race main.go
