package transform

import (
	"fmt"

	"github.com/kevin8428/factory_example/model"
	"github.com/kevin8428/factory_example/utils"
)

type CitiesTransform struct{}

var citiesFilters = []string{"international", "country", "continent"}

func (t *CitiesTransform) Get(s string) ([]model.Result, error) {
	return []model.Result{
		{
			Name: "Boston",
			Data: map[string]interface{}{
				"international": false,
				"country":       "USA",
				"continent":     "North America",
			},
		},
		{
			Name: "Paris",
			Data: map[string]interface{}{
				"international": false,
				"country":       "France",
				"continent":     "North America",
			},
		},
		{
			Name: "Boise",
			Data: map[string]interface{}{
				"international": false,
				"country":       "USA",
				"continent":     "North America",
			},
		},
	}, nil
}

func (t *CitiesTransform) Filter(s []model.Result, f string) ([]model.Result, error) {
	return []model.Result{}, nil
}

func (t *CitiesTransform) Append(s string) error {
	return nil
}

func (t *CitiesTransform) ShouldFilter(entity, filter string) (bool, error) {
	fmt.Println("should filter string param: ", filter)
	return utils.GetSliceMatch(filter, citiesFilters), nil
}
