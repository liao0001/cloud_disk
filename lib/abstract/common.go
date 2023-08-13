package abstract

// include
func strArrayIncludes(strA []string, s string) bool {
	return strArrayIndexOf(strA, s) != -1
}

func strArrayIndexOf(strA []string, s string) int {
	if strA == nil {
		return -1
	}
	for k, ts := range strA {
		if ts == s {
			return k
		}
	}
	return -1
}

func appendUnique(arr []string, s string) []string {
	if !strArrayIncludes(arr, s) {
		return append(arr, s)
	}
	return arr
}
