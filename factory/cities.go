package factory

import (
	"net/http"
)

type CitiesHandler struct {
	Server *http.ServeMux
}

func NewCitiesHandler(server *http.ServeMux) EntityHandlerInterface {
	return &CitiesHandler{
		Server: server,
	}
}

func (c *CitiesHandler) Process(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hit the cities handler"))
}
