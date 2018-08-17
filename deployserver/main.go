package main

import (
	"io"
	"log"
	"net/http"
	"os/exec"
)

func reLaunch() {
	cmd := exec.Command("sh", "./deploy.sh")
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	cmd.Wait()
}

func appHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "<h1>The server is relauching!</h1>")
	reLaunch()
}

func main() {
	http.HandleFunc("/", appHandler)
	err := http.ListenAndServe(":1234", nil)
	if err != nil {
		log.Fatal(err)
	}
}
