package factory

import (
	"net/http"
)

type AnimalsHandler struct {
	Server *http.ServeMux
}

func NewAnimalsHandler(server *http.ServeMux) EntityHandlerInterface {
	return &AnimalsHandler{
		Server: server,
	}
}

func (c *AnimalsHandler) Process(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hit the animals handler"))
}
