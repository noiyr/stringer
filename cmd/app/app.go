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



func(a* Application)InitConfig() *Application{

cfg:=&config.Config{
	Port: 4040,
	Env: "Prod",
}

a = &Application{
Config: cfg,
}


return a
}