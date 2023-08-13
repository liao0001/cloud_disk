package storage

import "io"

// 存储介质
type IStorage interface {
	// 推送文件
	PushFile(objName string, reader io.Reader) (link string, err error)
	// 添加目录
	CreateDir(dirName string) error
	// 删除文件
	DeleteFile(objName string) error
	// 获取文件链接
	GetUrl(objName string) (link string, err error)
	GetUrls(objNames []string) (links []string, err error)
	// 获取缩略图地址
	GetThumbUrl(objName string) (link string, err error)
	GetThumbUrls(objNames []string) (links []string, err error)
	// 获取文件
	GetFile(objName string) (io.Reader, error)
	GetFileTo(objName string, w io.Writer) error
	// 递归扫描
	ScanDst(dst string) ([]map[string]string, error)
}

// 配置
type Config struct {
	Thumb          ThumbConfig   `json:"thumb" yaml:"thumb"`                     // 缩略图配置
	DefaultStorage ConfigStorage `json:"default_storage" yaml:"default_storage"` // 默认存储引擎
	// OtherStorages []StorageConfig `json:"other_storages" yaml:"other_storages"` //后期可支持多个配置
}

type ConfigStorage struct {
	Key             string `json:"key" yaml:"key"`                             // 注册时用的key
	Driver          string `json:"driver" yaml:"driver"`                       // 类型
	AccessKeyID     string `json:"access_key_id" yaml:"access_key_id"`         //
	AccessKeySecret string `json:"access_key_secret" yaml:"access_key_secret"` //
	Endpoint        string `json:"endpoint" yaml:"endpoint"`                   // endpoint
	Bucket          string `json:"bucket" yaml:"bucket"`                       // bucket
	WithAcl         bool   `json:"with_acl" yaml:"with_acl"`
	Expiration      int    `json:"expiration" yaml:"expiration"` // 过期时间 单位:秒
}

type ThumbConfig struct {
	Enable bool   `json:"enable" yaml:"enable"`
	Width  int    `json:"width" yaml:"width"`
	Height int    `json:"height" yaml:"height"`
	FFMpeg string `json:"ffmpeg" yaml:"ffmpeg"`
}

type NewStorageFunc func(conf ConfigStorage) (IStorage, error)
