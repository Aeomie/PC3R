package main

import (
	"fmt"
	"net/http"
	"strings"
)

type DataStore interface {
	accCreation(username string, password int)
	getPw(username string) int
}

type Server struct {
	store DataStore
}

func (p *Server) ServeHTTP(w http.ResponseWriter,
	r *http.Request) {
	url := strings.TrimPrefix(r.URL.Path, "/accs/")
	switch r.Method {
	case http.MethodPost:
		p.handlePost(w, url)
	case http.MethodGet:
		p.handleGet(w, url)
	}
}

func (p *Server) handleGet(w http.ResponseWriter,
	url string) {
	fmt.Println("get")
}

func (p *Server) handlePost(w http.ResponseWriter,
	url string) {
	fmt.Println("post")
}
