package hakyhandlers



import(
"net/http"
)


type Routes struct{
	Mux *http.ServeMux
	Handler *Handlers
}


func(m *Routes)SetupRoutes() *Routes{

	m.Mux = http.NewServeMux()
	m.Mux.HandleFunc("v1/healthcheck",m.Handler.Healthcheckhandler)

return m
}