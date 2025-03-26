package main

import (
	"flag"
	"os"

	"github.com/nadiannis/claris/internal/log"
)

type config struct {
	port int
	env  string
}

type application struct {
	config config
	logger *log.Logger
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 8080, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|production)")

	flag.Parse()

	app := &application{
		config: cfg,
		logger: log.New(os.Stdout, log.LevelInfo),
	}

	err := app.serve()
	if err != nil {
		app.logger.LogFatal(err, nil)
	}
}
