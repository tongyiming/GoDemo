package main

import (
	"net/http"
	"fmt"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,"https service with one verify")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServeTLS(":8080",
		"src/https_verify_one/server.crt", "src/https_verify_one/server.key", nil)
}
