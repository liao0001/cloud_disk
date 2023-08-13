package qr_decode

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"gitee.com/monkey0001/my_tools/lib/utils"
	"os"
	"path/filepath"
	"strings"
)

var pluginPath string

func init() {
	pluginPath = `/Users/liaoyong/workspace/python/qrdecode/stable/qrdecode/qrdecode`
}

func SetPluginPath(p string) {
	pluginPath = p
}

func Decode(qrPath string) (string, error) {
	res, err := utils.ExecStd(pluginPath, qrPath)
	if err != nil {
		return "", err
	}
	//
	if len(res) == 0 {
		return "", fmt.Errorf("解析失败")
	}
	return strings.ReplaceAll(string(res), "\n", ""), nil
}

func DecodeUrl(url string) (string, error) {
	bs, err := utils.Download(url)
	if err != nil {
		return "", fmt.Errorf("下载文件失败:%s", err.Error())
	}
	fn := fmt.Sprintf("%x", md5.Sum(bs))

	filename := filepath.Join(os.TempDir(), fn)
	// 放入临时文件夹
	err = os.WriteFile(filename, bs, 0655)
	if err != nil {
		return "", fmt.Errorf("保存文件失败:%s", err.Error())
	}
	res, err := utils.ExecStd(pluginPath, filename)
	if err != nil {
		return "", err
	}
	//
	if len(res) == 0 {
		return "", fmt.Errorf("解析失败")
	}
	return strings.ReplaceAll(string(res), "\n", ""), nil
}

func DecodeBase64(val string) (string, error) {
	ss := strings.Split(val, ",")
	val = ss[len(ss)-1]

	bs, err := base64.StdEncoding.DecodeString(strings.Replace(val, " ", "+", -1))
	if err != nil {
		return "", fmt.Errorf("解析base64字符串失败:%s", err.Error())
	}
	fn := fmt.Sprintf("%x", md5.Sum(bs))

	filename := filepath.Join(os.TempDir(), fn)
	// 放入临时文件夹
	err = os.WriteFile(filename, bs, 0655)
	if err != nil {
		return "", fmt.Errorf("保存文件失败:%s", err.Error())
	}
	res, err := utils.ExecStd(pluginPath, filename)
	if err != nil {
		return "", err
	}
	//
	if len(res) == 0 {
		return "", fmt.Errorf("解析失败")
	}
	return strings.ReplaceAll(string(res), "\n", ""), nil
}
