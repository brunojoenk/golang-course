package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)
	log.Println("Executando...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

//~/dev/learning/golang/http/static(master)$ go run static.go
// Não roda pelo VS
