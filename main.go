package main

import (
	"fmt"
	"net/http"
	"text/template"
)

const portNumber = ":8080"

// Home is the handler for the home page
func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.tmpl")
}

// About is the handler for the about page
func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.tmpl")
}

// renderTemplate read, parse and render the template
func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error rendering template: ", err)
	}
}

// main is the main entry point
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Printf("Starting application on port %s\n", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
