package render

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}

// RenderTemplate reads, parses and renders the template using html/template
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	_, err := RenderTemplateTest()
	if err != nil {
		fmt.Println("Error gettign template cache", err)
	}

	parsedTemplate, _ := template.ParseFiles("templates/" + tmpl)

	err = parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error rendering template: ", err)
	}
}

func RenderTemplateTest() (map[string]*template.Template, error) {
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
