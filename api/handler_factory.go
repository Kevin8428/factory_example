package api

import (
	"fmt"
	"net/http"

	"github.com/kevin8428/factory_example/factory"
)

func (a *API) RegisterHandlers(server *http.ServeMux) {
	entities := []string{
		"cities",
		"animals",
		"people",
	}
	a.RegisterEntities(entities, server)
}

func (a *API) RegisterEntities(entities []string, server *http.ServeMux) {
	a.Handlers = map[string]factory.EntityHandlerInterface{}
	for _, entity := range entities {
		if factory := factory.EntityHandlerFactoryMethods[entity]; factory != nil {
			fmt.Println("registering handler for entity ", entity)
			handler := factory(server)
			a.Handlers[entity] = handler
		}
	}
}
