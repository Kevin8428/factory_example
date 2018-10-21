package utils

func GetSliceMatch(s string, sl []string) bool {
	for i := range sl {
		if sl[i] == s {
			return true
		}
	}
	return false
}
