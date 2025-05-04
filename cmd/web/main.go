package main

import (
	"flag"
	"html/template"
	"log/slog"
	"os"
	"sync"
)

type config struct {
	port int
	env  string
}

type application struct {
	config        config
	logger        *slog.Logger
	templateCache map[string]*template.Template
	wg            sync.WaitGroup
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	templateCache, err := newTemplateCache()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	app := &application{
		config:        cfg,
		templateCache: templateCache,
		logger:        logger,
	}

	err = app.serve()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}
