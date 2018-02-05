package main

import (
	"net/http"
	"fmt"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "https service demo")
}

func main() {
	http.HandleFunc("/", handler)
	//if is http:http.ListenAndServe(":8080", nil)
	http.ListenAndServeTLS(":8080", "src/https/server.crt", "src/https/server.key", nil)
}