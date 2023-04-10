package app

import (
	"hakydll/config"
	"hakydll/hakyhandlers"
	"hakydll/hakylog"
	"net/http"
)

type Application struct {
	Routes *hakyhandlers.Routes
	Log    *hakylog.Logging
	Config *config.Config
}

func (a *Application) InitConfig() *Application {

	cfg := &config.Config{
		Port: ":4040",
		Env:  "Prod",
	}

	handler := &hakyhandlers.Handlers{
		Config: cfg,
	}

	routes := &hakyhandlers.Routes{
		Mux:     http.NewServeMux(),
		Handler: handler,
	}

	a = &Application{
		Config: cfg,
		Routes: routes,
	}

	return a
}
