package encrypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"fmt"
)

/*
	双向加密
*/

//高级加密标准（Adevanced Encryption Standard ,AES）
//16,24,32位字符串的话，分别对应AES-128，AES-192，AES-256 加密方法

type AESService struct {
	secretKey []byte //验证秘钥
}

func NewAESService(secret string) (*AESService, error) {
	switch len(secret) {
	case 16, 24, 32:
		return &AESService{
			secretKey: []byte(secret),
		}, nil
	}
	return nil, fmt.Errorf("无效的秘钥长度(16,24,32):%d", len(secret))
}

//PKCS7 填充模式
func (s *AESService) pKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	//Repeat()函数的功能是把切片[]byte{byte(padding)}复制padding个，然后合并成新的字节切片返回
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

//填充的反向操作，删除填充字符串
func (s *AESService) pKCS7UnPadding(origData []byte) ([]byte, error) {
	//获取数据长度
	length := len(origData)
	if length == 0 {
		return nil, errors.New("加密字符串错误！")
	} else {
		//获取填充字符串长度
		unpadding := int(origData[length-1])
		//截取切片，删除填充字节，并且返回明文
		return origData[:(length - unpadding)], nil
	}
}

//实现加密
func (s *AESService) aesEcrypt(origData []byte) ([]byte, error) {
	//创建加密算法实例
	block, err := aes.NewCipher(s.secretKey)
	if err != nil {
		return nil, err
	}
	//获取块的大小
	blockSize := block.BlockSize()
	//对数据进行填充，让数据长度满足需求
	origData = s.pKCS7Padding(origData, blockSize)
	//采用AES加密方法中CBC加密模式
	blocMode := cipher.NewCBCEncrypter(block, s.secretKey[:blockSize])
	crypted := make([]byte, len(origData))
	//执行加密
	blocMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

//实现解密
func (s *AESService) aesDeCrypt(cypted []byte) ([]byte, error) {
	//创建加密算法实例
	block, err := aes.NewCipher(s.secretKey)
	if err != nil {
		return nil, err
	}
	//获取块大小
	blockSize := block.BlockSize()
	//创建加密客户端实例
	blockMode := cipher.NewCBCDecrypter(block, s.secretKey[:blockSize])
	origData := make([]byte, len(cypted))
	//这个函数也可以用来解密
	blockMode.CryptBlocks(origData, cypted)
	//去除填充字符串
	origData, err = s.pKCS7UnPadding(origData)
	if err != nil {
		return nil, err
	}
	return origData, err
}

//加密base64
func (s *AESService) EnCode(srt string) (string, error) {
	result, err := s.aesEcrypt([]byte(srt))
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(result), err
}

//解密
func (s *AESService) DeCode(tar string) (string, error) {
	//解密base64字符串
	pwdByte, err := base64.StdEncoding.DecodeString(tar)
	if err != nil {
		return "", err
	}
	//执行AES解密
	bs, err := s.aesDeCrypt(pwdByte)
	if err != nil {
		return "", err
	}
	return string(bs), nil
}
