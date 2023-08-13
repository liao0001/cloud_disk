package utils

import "strconv"

/*字符串和数字的互转*/

//字符串转为int64
func StrToInt64(s string, def ...int64) int64 {
	d := int64(0)
	if len(def) > 0 {
		d = def[0]
	}
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return d
	}
	return i
}

//字符串转为int
func StrToInt(s string, def ...int) int {
	d := 0
	if len(def) > 0 {
		d = def[0]
	}
	i, err := strconv.Atoi(s)
	if err != nil {
		return d
	}
	return i
}

//字符串转为float64
func StrToFloat64(s string, def ...float64) float64 {
	d := float64(0)
	if len(def) > 0 {
		d = def[0]
	}
	i, err := strconv.ParseFloat(s, 10)
	if err != nil {
		return d
	}
	return i
}
