package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Giuco/go-web/pkg/config"
	"github.com/Giuco/go-web/pkg/handlers"
	"github.com/Giuco/go-web/pkg/render"
)

const portNumber = ":8080"

// main is the main entry point
func main() {
	var app config.AppConfig
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("could not load template cache", err)
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Printf("Starting application on port %s\n", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
