package utils

// 字符串数组去重去空
func StrArrayDeal(strA []string) []string {
	if len(strA) > 0 {
		var newStrA []string
		for _, oa := range strA {
			ok := false
			if len(newStrA) > 0 {
				for _, na := range newStrA {
					if oa == na {
						ok = true
						break
					}
				}
			}
			if !ok && oa != "" {
				newStrA = append(newStrA, oa)
			}
		}
		strA = newStrA
	}
	return strA
}

// 字符串数组去重
func StrArrayUnique(strA []string) []string {
	if len(strA) > 0 {
		var newStrA []string
		for _, oa := range strA {
			exist := false
			if len(newStrA) > 0 {
				for _, na := range newStrA {
					if oa == na {
						exist = true
						break
					}
				}
			}
			if !exist {
				newStrA = append(newStrA, oa)
			}
		}
		strA = newStrA
	}
	return strA
}

// include
func StrArrayIncludes(strA []string, s string) bool {
	return StrArrayIndexOf(strA, s) != -1
}

func StrArrayIndexOf(strA []string, s string) int {
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

func AppendUnique(arr []string, s string) []string {
	if !StrArrayIncludes(arr, s) {
		return append(arr, s)
	}
	return arr
}

func AppendUniqueBatch(arr []string, ss ...string) []string {
	for _, s := range ss {
		if !StrArrayIncludes(arr, s) {
			arr = append(arr, s)
		}
	}
	return arr
}

func Union(slice1, slice2 []string) []string {
	m := make(map[string]int)
	for _, v := range slice1 {
		m[v]++
	}

	for _, v := range slice2 {
		times, _ := m[v]
		if times == 0 {
			slice1 = append(slice1, v)
		}
	}
	return slice1
}

//求交集
func Intersect(slice1, slice2 []string) []string {
	m := make(map[string]int)
	nn := make([]string, 0)
	for _, v := range slice1 {
		m[v]++
	}

	for _, v := range slice2 {
		times, _ := m[v]
		if times == 1 {
			nn = append(nn, v)
		}
	}
	return nn
}

//求差集 slice1-并集
func Difference(slice1, slice2 []string) []string {
	m := make(map[string]int)
	nn := make([]string, 0)
	inter := Intersect(slice1, slice2)
	for _, v := range inter {
		m[v]++
	}

	for _, value := range slice1 {
		times, _ := m[value]
		if times == 0 {
			nn = append(nn, value)
		}
	}
	return nn
}
