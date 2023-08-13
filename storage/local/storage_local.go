package local

import (
	"fmt"
	"github.com/liao0001/cloud_disk/storage"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path/filepath"
	"strings"
)

const defaultDstPath = "dst_files"
const defaultFileName = "default"

type Local struct {
	dstPath    string
	httpPrefix string
	withAcl    bool
}

func NewStorage(conf storage.ConfigStorage) (storage.IStorage, error) {
	httpPrefix := conf.Endpoint
	dstPath := conf.Bucket

	var err error
	if len(dstPath) == 0 {
		dstPath = defaultDstPath
	}
	dstPath, err = filepath.Abs(dstPath)
	if err != nil {
		logrus.Fatalf("初始化路径失败:%s", err.Error())
	}
	fileInfo, err := os.Stat(dstPath)
	if err != nil {
		_ = os.MkdirAll(dstPath, 0755)
	}
	if !fileInfo.IsDir() {
		logrus.Fatalf("目标路径已存在，但不是个目录:%s", dstPath)
	}
	httpPrefix = strings.TrimRight(httpPrefix, "/")
	httpPrefix = httpPrefix + "/"
	//
	_ = os.Chmod(dstPath, 0755)

	serv := &Local{
		dstPath:    dstPath,
		httpPrefix: httpPrefix,
		withAcl:    conf.WithAcl,
	}
	return serv, err
}

func (s *Local) PushFile(objName string, reader io.Reader) (link string, err error) {
	if len(objName) == 0 {
		objName = defaultFileName
	}

	err = s.createDirByObjName(objName)
	if err != nil {
		return "", errors.Wrap(err, "创建文件夹失败")
	}

	wholePath := s.getAbsPath(objName)
	file, err := os.Create(wholePath)
	if err != nil {
		return "", err
	}
	defer func() {
		_ = file.Close()
	}()

	bufferBs := make([]byte, 1024*1024)

	_, err = io.CopyBuffer(file, reader, bufferBs)
	if err != nil {
		return "", err
	}
	// fmt.Println("总大小:", size)
	return s.GetUrl(objName)
}

func (s *Local) CreateDir(dirName string) error {
	wholePath := s.getAbsPath(dirName)
	info, err := os.Stat(wholePath)
	if err != nil {
		return os.MkdirAll(wholePath, 0755)
	}
	if !info.IsDir() {
		return fmt.Errorf("路径获取失败,已存在同名文件:%s ", dirName)
	}
	return nil
}

// 仅删除文件 目录保留
func (s *Local) DeleteFile(objName string) error {
	return os.Remove(s.getAbsPath(objName))
}

func (s *Local) GetUrl(objName string) (link string, err error) {
	objName = s.resetObjName(objName)
	return s.getUrlByObjName(objName), nil
}

func (s *Local) GetUrls(objNames []string) (links []string, err error) {
	for _, one := range objNames {
		u, err := s.GetUrl(one)
		if err != nil {
			return nil, err
		}
		links = append(links, u)
	}
	return links, nil
}

func (s *Local) GetThumbUrl(objName string) (link string, err error) {
	objName = s.resetObjName(objName)

	thumbName := s.calcThumbUrl(objName)
	_, err = os.Stat(filepath.Join(s.dstPath, thumbName))
	if err != nil {
		return "", err
	}
	return s.getUrlByObjName(thumbName), nil
}

func (s *Local) GetThumbUrls(objNames []string) (links []string, err error) {
	// TODO implement me
	panic("implement me")
}

func (s *Local) GetFile(objName string) (io.Reader, error) {
	// TODO implement me
	panic("implement me")
}

func (s *Local) GetFileTo(objName string, w io.Writer) error {
	// TODO implement me
	panic("implement me")
}

func (s *Local) ScanDst(dst string) ([]map[string]string, error) {
	// TODO implement me
	panic("implement me")
}

// 创建文件夹
func (s *Local) createDirByObjName(objName string) error {
	wholePath := s.getAbsPath(objName)
	dir := filepath.Dir(wholePath)
	if dir == s.dstPath {
		return nil
	}
	info, err := os.Stat(dir)
	if err != nil {
		return os.MkdirAll(dir, 0755)
	}
	if !info.IsDir() {
		return fmt.Errorf("路径获取失败,已存在同名文件:%s ", objName)
	}
	_ = os.MkdirAll(dir, 0755)
	_ = os.Chmod(dir, 0755)
	return nil
}

// 根据objName 获取绝对路径
func (s *Local) getAbsPath(objName string) string {
	wholePath := filepath.Join(s.dstPath, s.resetObjName(objName))
	return wholePath
}

func (s *Local) resetObjName(objName string) string {
	// 去除特殊字符(不能当做文件名称的全改成_)
	objName = strings.ReplaceAll(objName, "\\", "/")
	objName = strings.ReplaceAll(objName, "/", string(os.PathSeparator))
	objName = strings.TrimLeft(objName, string(os.PathSeparator))
	return objName
}

func (s *Local) calcThumbUrl(objName string) string {
	ext := filepath.Ext(objName)
	if len(ext) == 0 {
		return objName + "_thumb"
	}
	return strings.Replace(objName, ext, "", len(objName)-1) + "_thumb" + ext
}

func (s *Local) getUrlByObjName(objName string) string {
	return s.httpPrefix + objName
}
