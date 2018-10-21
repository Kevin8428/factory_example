package transform

import (
	"errors"
	"fmt"
)

type TransformerByEntity struct {
	Transformer map[string]Transformer
}

type Transformer interface {
	Get(s string) ([]string, error)
	Filter(s []string, f string) ([]string, error)
	Append(s string) error
	ShouldFilter(s string) bool
}

func (t *TransformerByEntity) Get(s string) ([]string, error) {
	transformer, ok := t.Transformer[s]

	if !ok {
		fmt.Println("no transform entity exists for ", s)
		return nil, errors.New("no transform entity exists")
	}
	return transformer.Get(s)
}

func (t *TransformerByEntity) Filter(s []string, f string) ([]string, error) {
	return []string{}, nil
}

func (t *TransformerByEntity) Append(s string) error {
	return nil
}

func (t *TransformerByEntity) ShouldFilter(s string) bool {
	transformer, ok := t.Transformer[s]
	if !ok {
		fmt.Println("no transform entity exists for ", s)
	}
	return transformer.ShouldFilter(s)
}
