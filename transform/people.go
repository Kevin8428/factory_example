package transform

import "github.com/kevin8428/factory_example/model"

type PeopleTransform struct{}

func (t *PeopleTransform) Get(s string) ([]model.Result, error) {
	return []model.Result{
		{
			Name: "Kevin",
			Data: map[string]interface{}{
				"gender":   "male",
				"countryd": "USA",
			},
		},
		{
			Name: "Greg",
			Data: map[string]interface{}{
				"gender":   "male",
				"countryd": "USA",
			},
		},
		{
			Name: "Elyse",
			Data: map[string]interface{}{
				"gender":  "female",
				"country": "switzerland",
			},
		},
	}, nil
}

func (t *PeopleTransform) Filter(s string, r []model.Result, f string) ([]model.Result, error) {
	return []model.Result{}, nil
}
func (t *PeopleTransform) Append(s string) error {
	return nil
}

func (t *PeopleTransform) ShouldFilter(entity, filter string) (bool, error) {
	return false, nil
}
