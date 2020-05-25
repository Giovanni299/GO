package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	startWatch := time.Now()
	channel1 := make(chan string)
	//Lista de servidores.
	servers := []string{
		"http://platzi.com",
		"http://google.com",
		"http://facebook.com",
		"http://instagram.com",
		"http://pruebaInvalida123.com",
	}

	//Recorrer la lista de servidores y generar un hilo por cada una.
	for _, server := range servers {
		go checkStatusServer(server, channel1)
	}

	//Esperar la respuesta de todos los Goroutines.
	for index := range servers {
		fmt.Println(index+1, <-channel1)
	}

	//Mostar tiempo de duración de la ejecución.
	timeExec := time.Since(startWatch)
	fmt.Printf("Execution time: %s", timeExec)
}

//Revisar el estatus de cada uno de los servidores.
func checkStatusServer(server string, channel chan string) {
	_, err := http.Get(server)
	if err != nil {
		channel <- server + " Is not Available =("
	} else {
		channel <- server + " It's working OK"
	}
}
