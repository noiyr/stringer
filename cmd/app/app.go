package app

import (

	"hakydll/hakyhandlers"
	"hakydll/config"
	"hakydll/hakylog"
)

type Application struct {
	Routes	*hakyhandlers.Routes
	Log    *hakylog.Logging
	Config *config.Config
}



func(a* Application)initApp() (app *Application, err error) {





return a,err
}