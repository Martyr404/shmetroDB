package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func CreateH2Server() {
	path, _ := os.Getwd()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello world")
		fmt.Fprintf(w, "Protocol: %s\n", r.Proto)
	})
	h2s := &http2.Server{
		// ...
	}
	h1s := &http.Server{
		Addr:    ":9988",
		Handler: h2c.NewHandler(handler, h2s),
	}
	log.Fatal(h1s.ListenAndServeTLS(path+"\\certificate.pem", path+"\\private.pem"))
}
