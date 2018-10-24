package utils

import "github.com/kevin8428/factory_example/model"

func GetSliceMatch(s string, sl []string) bool {
	for i := range sl {
		if sl[i] == s {
			return true
		}
	}
	return false
}

func FilterResults(r []model.Result, filter string) bool {

	return false
}
