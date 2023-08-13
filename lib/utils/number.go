package utils

import (
	"fmt"
	"math"
	"strings"
)

//数字保留对应小数(默认6位)
func KeepFloat(f float64, digest ...int) float64 {
	n := 6
	if len(digest) > 0 && digest[0] >= 0 {
		n = digest[0]
	}
	return round(f, n)
}

func round(f float64, n int) float64 {
	n10 := math.Pow10(n)
	if f >= 0 {
		return math.Trunc((f+0.5/n10)*n10) / n10
	}
	return math.Trunc((f-0.5/n10)*n10) / n10
}

//除
func Divide(f float64, nf float64, digest ...int) float64 {
	if nf == 0 {
		return 0
	}
	return KeepFloat(f/nf, digest...)
}

//求和
func Sum(fs []float64, digest ...int) float64 {
	if len(fs) == 0 {
		return 0
	}
	total := float64(0)
	for _, f := range fs {
		total += f
	}
	return KeepFloat(total, digest...)
}

//平均值
func Average(fs []float64, digest ...int) float64 {
	if len(fs) == 0 {
		return 0
	}
	total := float64(0)
	for _, f := range fs {
		total += f
	}

	return Divide(total, float64(len(fs)), digest...)
}

//最大值
func Max(fs []float64, digest ...int) float64 {
	if len(fs) == 0 {
		return 0
	}
	max := fs[0]
	for _, f := range fs {
		if f > max {
			max = f
		}
	}

	return KeepFloat(max, digest...)
}

//最小值
func Min(fs []float64, digest ...int) float64 {
	if len(fs) == 0 {
		return 0
	}
	min := fs[0]
	for _, f := range fs {
		if f < min {
			min = f
		}
	}

	return KeepFloat(min, digest...)
}

//避免科学计数法的数字转文本
func NumberToString(f float64) string {
	if f == 0 {
		return "0"
	}
	return strings.TrimRight(fmt.Sprintf("%f", KeepFloat(f)), "0")
}
