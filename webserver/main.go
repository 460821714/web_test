package main

import (
	"io"
	"log"
	"net/http"
)

func appHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "<h1>hello,有一点点奇怪.name is Jake.add Testing.what happened!</h1>")
}

func main() {
	http.HandleFunc("/", appHandler)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
