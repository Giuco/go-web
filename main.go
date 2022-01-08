package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// Home is the handler for the home page
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This the home page")
}

// About is the handler for the about page
func About(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the about page. And 2+2=%d", addValues(2, 2))
}

// addValues sums two values
func addValues(a, b int) int {
	return a + b
}

// main is the main entry point
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Printf("Starting application on port %s\n", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
