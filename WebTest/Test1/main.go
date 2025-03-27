package main

import (
	"html/template"
	"net/http"
)

// Define a struct to hold the data
type PageData struct {
	Title string
	Users []string
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Prepare dynamic data to be passed into the template
	data := PageData{
		Title: "User List",
		Users: []string{"Alice", "Bob", "Charlie"},
	}

	// Parse the template file
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		return
	}

	// Execute the template and pass the dynamic data
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
		return
	}
}

func main() {
	// Serve static files from the "static" directory
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Handle the root route
	http.HandleFunc("/", handler)

	// Start the server on port 8080
	http.ListenAndServe(":8080", nil)
}
