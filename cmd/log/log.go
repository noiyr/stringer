package hakylog

import (
	"hakydll/config"
	"log"
	"os"
)

type Logging struct {
	Config *config.Config
	Logger *log.Logger
}

func (l *Logging) LogInit() (err error) {

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	l.Logger = logger


	return err
}
