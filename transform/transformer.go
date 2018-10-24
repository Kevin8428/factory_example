package transform

import (
	"errors"

	"github.com/kevin8428/factory_example/model"
)

type TransformerByEntity struct {
	Transformer map[string]Transformer
}

type Transformer interface {
	Get(s string) ([]model.Result, error)
	Filter(s []model.Result, f string) ([]model.Result, error)
	Append(s string) error
	ShouldFilter(entity, filter string) (bool, error)
}

func (t *TransformerByEntity) Get(s string) ([]model.Result, error) {
	transformer, ok := t.Transformer[s]

	if !ok {
		return nil, errors.New("no transform entity exists for entity: " + s)
	}
	return transformer.Get(s)
}

func (t *TransformerByEntity) Filter(s []model.Result, f string) ([]model.Result, error) {
	return []model.Result{}, nil
}

func (t *TransformerByEntity) Append(s string) error {
	return nil
}

func (t *TransformerByEntity) ShouldFilter(entity, filter string) (bool, error) {
	transformer, ok := t.Transformer[entity]
	if !ok {
		return false, errors.New("no transform entity exists")
	}
	return transformer.ShouldFilter(entity, filter)
}
