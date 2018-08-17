package main

import (
	"io"
	"log"
	"net/http"
)

func appHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "<h1>hello,my name is Jake.add Testing.</h1>")
}

func main() {
	http.HandleFunc("/", appHandler)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
