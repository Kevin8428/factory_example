package transform

type PeopleTransform struct{}

func (t *PeopleTransform) Get(s string) ([]string, error) {
	return []string{"Kevin", "Greg", "Elyse", "Klaus", "Julia"}, nil
}

func (t *PeopleTransform) Filter(s []string, f string) ([]string, error) {
	return []string{}, nil
}
func (t *PeopleTransform) Append(s string) error {
	return nil
}

func (t *PeopleTransform) ShouldFilter(s string) bool {
	return false
}
