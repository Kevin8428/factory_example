package transform

import "github.com/kevin8428/factory_example/utils"

type CitiesTransform struct{}

var citiesFilters = []string{"international", "country", "continent"}

func (t *CitiesTransform) Get(s string) ([]string, error) {
	return []string{"Boston", "Paris", "San Francisco", "New York", "Munich"}, nil
}

func (t *CitiesTransform) Filter(s []string, f string) ([]string, error) {
	return []string{}, nil
}

func (t *CitiesTransform) Append(s string) error {
	return nil
}

func (t *CitiesTransform) ShouldFilter(s string) bool {
	return utils.GetSliceMatch(s, citiesFilters)
}
