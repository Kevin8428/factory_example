package utils

import (
	"github.com/kevin8428/factory_example/model"
)

func GetSliceMatch(s string, sl []string) bool {
	for i := range sl {
		if sl[i] == s {
			return true
		}
	}
	return false
}

func FilterResults(r []model.Result, filter string) ([]model.Result, error) {
	filtered := []model.Result{}
	for _, v := range r {
		switch v.Data[filter].(type) {
		case bool:
			if b, ok := v.Data[filter].(bool); ok {
				if b {
					filtered = append(filtered, v)
				}
			}
		case string:

		case int:

		}
	}
	return filtered, nil
}
