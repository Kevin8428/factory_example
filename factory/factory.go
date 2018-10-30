package factory

import "net/http"

var EntityHandlerFactoryMethods = map[string]EntityHandlerFactory{
	"cities":  NewCitiesHandler,
	"people":  NewPeopleHandler,
	"animals": NewAnimalsHandler,
}

type EntityHandlerFactory func(server *http.ServeMux) EntityHandlerInterface

type EntityHandlerInterface interface {
	Process(w http.ResponseWriter, r *http.Request)
}
