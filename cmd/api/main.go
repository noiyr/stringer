package main

import(
	"fmt"
	"flag"
	"net/http"
	"os"
	"time"
	"log"
)

const(
version = 1.1

)

type Application struct{
	log* log.Logger
	config Config
}

type Config struct{
	port int
	env string
}


func main(){
var cfg Config

flag.IntVar(&cfg.port,"port",4000,"use in port http")
flag.StringVar(&cfg.env,"enviromnent","Environment (development|staging|production","development")
flag.Parse()


logger:=log.New(os.Stdout,"",log.Ldate|log.Ltime)
app:=&Application{
	config: cfg,
	log: logger,
}

mux:= http.NewServeMux()
mux.HandleFunc("/v1/healthcheck",app.healthcheckhandler)



srv:= &http.Server{
	Addr: fmt.Sprintf(":%d",cfg.port),
	Handler: mux,
	IdleTimeout: time.Minute,
	ReadTimeout: 10 * time.Second,
	WriteTimeout: 30* time.Second,
}

logger.Printf("starting server %s on %s port",app.config.env,app.config.port)
err:= srv.ListenAndServe()
logger.Fatal(err)
}