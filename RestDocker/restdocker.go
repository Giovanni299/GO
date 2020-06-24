package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello %s!", r.URL.Path[1:])
	fmt.Println("RESTfulServ. on:8093, Controller:", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Starting Restful services...")
	fmt.Println(GetOutboundIP())
	fmt.Println("Using port:8093")
	err := http.ListenAndServe(":8093", nil)
	log.Print(err)
	errorHandler(err)
}

func errorHandler(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

/*
	Para ejecutarlo:
	go run restdocker.go
	1. docker build -t restdocker .   //Contruir la imagen
	2. docker run -p 8093:8093 restdocker    //Ejecutar la aplicación
	3. ir  la URl http://localhost:8093/qwe
*/
