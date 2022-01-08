package main

import (
	"errors"
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

// Divide is the handler for the divide page
func Divide(w http.ResponseWriter, r *http.Request) {
	a, b := 100., 0.
	v, err := divideValues(a, b)
	if err != nil {
		fmt.Fprintf(w, "Cannot divide by 0")
		return
	}
	fmt.Fprintf(w, "%v / % v = %v", a, b, v)
}

// divideValues divides two values
func divideValues(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("Cannot divide by 0")
	}
	return a / b, nil
}

// addValues sums two values
func addValues(a, b int) int {
	return a + b
}

// main is the main entry point
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	fmt.Printf("Starting application on port %s\n", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
