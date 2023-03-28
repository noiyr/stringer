package main

import (
	"fmt"
	"net/http"
)



func(a* Application)healthcheckhandler(w http.ResponseWriter,r* http.Request){

	fmt.Fprintf(w,"%s or %q",a.config.env,a.config.port)

}