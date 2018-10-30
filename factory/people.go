package factory

import (
	"net/http"
)

type PeopleHandler struct {
	Server *http.ServeMux
}

func NewPeopleHandler(server *http.ServeMux) EntityHandlerInterface {
	return &PeopleHandler{
		Server: server,
	}
}

func (c *PeopleHandler) Process(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hit the people handler"))
}
