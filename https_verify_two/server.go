package main

import (
	"io/ioutil"
	"fmt"
	"net/http"
	"crypto/tls"
	"crypto/x509"
)

type myhandler struct {
}

func (h *myhandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "https service with two verify")
}

func main() {
	pool := x509.NewCertPool()
	caCertPath := "src/https_verify_two/ca.crt"

	caCrt, err := ioutil.ReadFile(caCertPath)
	if err != nil {
		fmt.Println("ReadFile err:", err)
		return
	}
	pool.AppendCertsFromPEM(caCrt)

	s := &http.Server{
		Addr:    ":8080",
		Handler: &myhandler{},
		TLSConfig: &tls.Config{
			ClientCAs:  pool,
			ClientAuth: tls.RequireAndVerifyClientCert,
		},
	}

	err = s.ListenAndServeTLS("src/https_verify_two/server.crt", "src/https_verify_two/server.key")
	if err != nil {
		fmt.Println("ListenAndServeTLS err:", err)
	}
}