package utils

import (
	"fmt"
	"testing"
	"time"
)

//字符串md5加密
func TestMD5(t *testing.T) {
	userName := "datav-1"
	password := "bepIStckBARI"
	tag := "202008050930329962127"
	s := fmt.Sprintf("%s%s%s%s", time.Now().Format("20060102"), userName, password, tag)
	s2 := fmt.Sprintf("%s%s%s", userName, password, tag)
	fmt.Println(MD5(s))
	fmt.Println(MD5(s2))
}
