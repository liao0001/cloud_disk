package config

import (
	"github.com/liao0001/cloud_disk/db"
	"github.com/liao0001/cloud_disk/storage"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"path/filepath"
)

// 配置文件
type Conf struct {
	// 版本
	Version string `yaml:"-"`
	// 配置文件目录
	ConfigDir string `yaml:"-"`
	// 运行模式
	Runmode string `yaml:"runmode"`
	LogDir  string `json:"log_dir" yaml:"log_dir"` // 日志地址
	Http    struct {
		Port    int `json:"port" yaml:"port"` // 端口
		Captcha struct {
			Enable   bool `yaml:"enable"`   // 是否需要验证码(仅密码登录有效)
			Length   int  `yaml:"length"`   // 长度 4-8  (默认 4)
			Duration int  `yaml:"duration"` // 有效期  单位秒 (默认 180)
		} `yaml:"captcha"` // 验证码配置
		TokenDuration        int    `yaml:"token_duration"`         // token有效期
		DefaultPasswordMode  string `yaml:"default_password_mode"`  // 默认密码模式  目前仅支持固定密码 (后续可以加上正则或字段值什么的)
		DefaultPasswordValue string `yaml:"default_password_value"` // 默认密码
	} `json:"http" yaml:"http"`
	DB      db.Config      `json:"db" yaml:"db"`
	Storage storage.Config `json:"storage" yaml:"storage"`
}

// 载入配置
func LoadConfig(path string) (*Conf, error) {
	bs, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var res Conf
	err = yaml.Unmarshal(bs, &res)
	if err != nil {
		return nil, err
	}

	res.ConfigDir = filepath.Dir(path)

	// 读取同目录下的version文件(若没有  版本为空)
	versionPath := filepath.Join(res.ConfigDir, "version")
	bs, _ = ioutil.ReadFile(versionPath)
	res.Version = string(bs)

	return &res, nil
}
