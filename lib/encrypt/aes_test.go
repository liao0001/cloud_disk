package encrypt

import (
	"fmt"
	"testing"
)

func TestErrorSecret(t *testing.T) {
	s, err := NewAESService("123456789")
	if err != nil {
		fmt.Println(err.Error())
	}
	s, err = NewAESService("1234567890123456")
	if err != nil {
		fmt.Println(err.Error())
	}
	s, err = NewAESService("1234567890123456/*-+.=-0")
	if err != nil {
		fmt.Println(err.Error())
	}
	s, err = NewAESService("1234567890123456/*-+.=-0+_)(*&^%")
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(s)
}

func TestEncrypt(t *testing.T) {
	s, err := NewAESService("1234567890123456/*-+.=-0+_)(*&^%")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// token := fmt.Sprintf("suodaoOM,%d,202204181729101774300000", time.Now().Add(7*24*time.Hour).Unix())
	token := "suodaoOM"
	fmt.Println("加密前长度：", len(token))
	enStr, err := s.EnCode(token)
	if err != nil {
		panic("加密失败:" + err.Error())
	}
	fmt.Println("加密结果：", enStr, "  ", len(enStr))

	res, err := s.DeCode(enStr)
	if err != nil {
		panic("解密失败:" + err.Error())
	}
	fmt.Println("解密结果： ", res == token)
}
