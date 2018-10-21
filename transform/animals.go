package transform

type AnimalsTransform struct{}

func (t *AnimalsTransform) Get(s string) ([]string, error) {
	return []string{"Dog", "Cat", "Cow", "Moose", "Horse"}, nil
}

func (t *AnimalsTransform) Filter(s []string, f string) ([]string, error) {
	return []string{}, nil
}
func (t *AnimalsTransform) Append(s string) error {
	return nil
}

func (t *AnimalsTransform) ShouldFilter(s string) bool {
	return false
}
