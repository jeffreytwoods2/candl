package main

import (
	"html/template"
	"io/fs"
	"path/filepath"

	"candl.jwoods.dev/ui"
)

type templateData struct {
	CurrentYear    int
	Ticker         string
	Timestamps     []string
	Closes         []float64
	Highs          []float64
	Lows           []float64
	Opens          []float64
	MaxPrice       float64
	MinPrice       float64
	FirstTimestamp string
	LastTimestamp  string
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := fs.Glob(ui.Files, "html/pages/*.tmpl")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		patterns := []string{
			"html/base.tmpl",
			"html/partials/*.tmpl",
			page,
		}

		ts, err := template.New(name).ParseFS(ui.Files, patterns...)
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil
}
