package main

import (
	"fmt"
	"net/http"

	"github.com/Giuco/go-web/pkg/handlers"
)

const portNumber = ":8080"

// main is the main entry point
func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Starting application on port %s\n", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
