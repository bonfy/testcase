package main

import (
	"fmt"
	"net/http"
)

func greetHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func main() {
	http.HandleFunc("/", greetHandler)
	http.ListenAndServe(":8080", nil)
}
