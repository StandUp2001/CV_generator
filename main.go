package main

import (
	"encoding/json"
	"html/template"
	"os"
	"path/filepath"
)

func main() {
	// Read JSON
	data, err := os.ReadFile("cv.json")
	if err != nil {
		panic(err)
	}

	var cv CV
	if err := json.Unmarshal(data, &cv); err != nil {
		panic(err)
	}

	init_()

	// Create HTML output
	f, err := os.Create("output/index.html")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	funcMap := template.FuncMap{
		"dots": func(n int) template.HTML {
			return template.HTML(dots(n))
		},
	}

	tmplPath := filepath.Join("template", "layout.html")
	tmpl := template.Must(template.New("cv").Funcs(funcMap).ParseFiles(tmplPath))
	err = tmpl.ExecuteTemplate(f, "layout.html", cv)
	if err != nil {
		panic(err)
	}

	println("CV (HTML) generated successfully in 'output' directory.")

}
