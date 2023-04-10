package main

import (
	_"fmt"
	"hakydll/app"
	"net/http"
)


func main(){

	haky:=app.Application{}
	cfg:= haky.InitConfig()


	mux:=cfg.Routes.Mux
	mux.HandleFunc("/v1/",cfg.Routes.Handler.Healthcheckhandler)
	
	http.ListenAndServe(cfg.Config.Port,mux)
}
















/* 


package main

import (
	"log"
)

type Config struct {
	AppName   string
	LogFile   string
	LogLevel  string
}

type Log struct {
	Config Config
	Logger *log.Logger
}

func (l *Log) setupLog() error {

	logger := log.New(file, l.Config.AppName+": ", log.LstdFlags)

	l.Logger = logger

	return nil
}

func main() {
	myLog := Log{
		Config: Config{
			AppName:  "MyApp",
			LogFile:  "/var/log/myapp.log",
			LogLevel: "info",
		},
	}

	err := myLog.setupLog()
	if err != nil {
		log.Fatal(err)
	}

	myLog.Logger.Println("This is a log message")

	}
}

*/