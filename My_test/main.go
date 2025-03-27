package main

import (
	"log"
	"net/http"
)

func main() {
	server := &Server{NewMemoryStore()}
	log.Fatal(http.ListenAndServe(":8000", server))
}
