package transform

import (
	"github.com/kevin8428/factory_example/model"
	"github.com/kevin8428/factory_example/utils"
)

type CitiesTransform struct{}

var citiesFilters = []string{"isInternational", "country", "continent"}

func (t *CitiesTransform) Get(s string) ([]model.Result, error) {
	return []model.Result{
		{
			Name: "Boston",
			Data: map[string]interface{}{
				"isInternational": false,
				"country":         "USA",
				"continent":       "North America",
			},
		},
		{
			Name: "Paris",
			Data: map[string]interface{}{
				"isInternational": true,
				"country":         "France",
				"continent":       "Europe",
			},
		},
		{
			Name: "Boise",
			Data: map[string]interface{}{
				"isInternational": false,
				"country":         "USA",
				"continent":       "North America",
			},
		},
	}, nil
}

func (t *CitiesTransform) Filter(s string, r []model.Result, f string) ([]model.Result, error) {
	return utils.FilterResults(r, f)
}

func (t *CitiesTransform) Append(s string) error {
	return nil
}

func (t *CitiesTransform) ShouldFilter(entity, filter string) (bool, error) {
	return utils.GetSliceMatch(filter, citiesFilters), nil
}
