package transform

import "github.com/kevin8428/factory_example/model"

type AnimalsTransform struct{}

func (t *AnimalsTransform) Get(s string) ([]model.Result, error) {
	return []model.Result{
		{
			Name: "Dog",
			Data: map[string]interface{}{
				"legs":         4,
				"domesticated": true,
				"canFly":       false,
			},
		},
		{
			Name: "Cat",
			Data: map[string]interface{}{
				"legs":         4,
				"domesticated": true,
				"canFly":       false,
			},
		},
		{
			Name: "Bird",
			Data: map[string]interface{}{
				"legs":         4,
				"domesticated": false,
				"canFly":       true,
			},
		},
	}, nil
}

func (t *AnimalsTransform) Filter(s string, r []model.Result, f string) ([]model.Result, error) {
	return []model.Result{}, nil
}
func (t *AnimalsTransform) Append(s string) error {
	return nil
}

func (t *AnimalsTransform) ShouldFilter(entity, filter string) (bool, error) {
	return false, nil
}
