package encrypt

import (
	"crypto/md5"
	"fmt"
)

/*
	单向加密
*/

//md5加密
func MD5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

//md5 16位秘钥
func MD5Short(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))[8:24]
}
