package main

import (
	"html/template"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Parse the base layout and the page template
	tmpl, err := template.New("base.html").ParseFiles("base.html", "home.html")
	if err != nil {
		http.Error(w, "Error parsing templates", http.StatusInternalServerError)
		return
	}

	// Render the template
	err = tmpl.ExecuteTemplate(w, "base.html", nil)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}

func main() {
	// Handle the root route
	http.HandleFunc("/", handler)

	// Start the server
	log.Fatal(http.ListenAndServe(":8080", nil))
}
