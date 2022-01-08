package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}

// RenderTemplate reads, parses and renders the template using html/template
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	tc, err := CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("template not found:", t)
	}

	buf := new(bytes.Buffer)
	err = t.Execute(buf, nil)
	if err != nil {
		log.Fatal("could not execute template:", err)
	}

	_, err = buf.WriteTo(w)
	if err != nil {
		log.Fatal("could not write template to response:", err)
	}

}

// CreateTemplateCache creates a template cache as a map
func CreateTemplateCache() (map[string]*template.Template, error) {
	templateCache := make(map[string]*template.Template)

	// this gets a list of all files ending with page.tmpl, and stores
	// it in a slice of strings called pagesPaths
	pagesPaths, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return templateCache, err
	}

	// this gets a list of all files ending with layout.tmpl, and stores
	// it in a slice of strings called layoutsPaths
	layoutsPaths, err := filepath.Glob("./templates/*.layout.tmpl")
	if err != nil {
		return templateCache, err
	}
	hasLayouts := len(layoutsPaths) > 0

	// now we loop through the slice of strings, which has two
	// entries: "home.page.tmpl" and "about.page.tmpl"
	for _, pagePath := range pagesPaths {
		// creates a new page
		name := filepath.Base(pagePath)
		fmt.Println("Page is currently", name)
		ts, err := template.New(name).Funcs(functions).ParseFiles(pagePath)
		if err != nil {
			return templateCache, err
		}

		// parses all the layouts
		if hasLayouts {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return templateCache, err
			}
		}

		// adds the page to templateCache
		templateCache[name] = ts
	}

	return templateCache, nil
}
