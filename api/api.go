package api

import "github.com/kevin8428/factory_example/transform"

type API struct {
	transformer transform.Transformer
}

func New() *API {
	return &API{
		transformer: &transform.TransformerByEntity{
			Transformer: map[string]transform.Transformer{
				"animals": &transform.AnimalsTransform{},
				"people":  &transform.PeopleTransform{},
				"cities":  &transform.CitiesTransform{},
			},
		},
	}
}
