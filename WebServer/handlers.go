package main

import (
	"fmt"
	"net/http"
)

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func HandleAPI(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the route to API.")
}
