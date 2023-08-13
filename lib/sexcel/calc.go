package sexcel

import (
	"fmt"
	"strings"
)

//根据列号和行号返回excel单元格标识  rowIndex会+1返回
func GetAxis(colIndex, rowIndex int) string {
	return fmt.Sprintf("%s%d", GetColAxis(colIndex), rowIndex+1)
}

// A=0; Z=25; AA=26 = 26*1+0;
func GetColAxis(colIndex int) string {
	//最多支持6个数 (6个数支持到3亿多了)
	letters := make([]string, 6)
	//位置计数 遍历每个位置的剩的数字
	lrem := colIndex
	var lmod int
	//字母位置
	posi := len(letters) - 1
	for ; lrem >= 0 && posi >= 0; lrem = lrem/26 - 1 {
		//对应倍数下的余数 就是指定位置上的数了
		lmod = lrem % 26
		letters[posi] = fmt.Sprintf("%c", 65+lmod)
		posi--
	}
	return strings.Join(letters, "")
}

//范围判断 这里-1表示不限
func InRangeWide(rangeStart, rangeEnd, target int) bool {
	return (target >= rangeStart || rangeStart == -1) && (target >= rangeEnd || rangeEnd == -1)
}
