package hakyhandlers

import (
	"fmt"
	"net/http"
	"hakydll/config"
)

type Handlers struct {
Config *config.Config
}

func (a Handlers) Healthcheckhandler(w http.ResponseWriter, r *http.Request) {
	
fmt.Fprintf(w,"Init at %s",a.Config.Env)
}
