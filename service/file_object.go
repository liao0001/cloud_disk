package service

import (
	"github.com/liao0001/cloud_disk"
	"github.com/liao0001/cloud_disk/db"
	"github.com/liao0001/cloud_disk/lib/utils"
	"github.com/liao0001/cloud_disk/models"
	"github.com/liao0001/cloud_disk/storage"
	"github.com/sirupsen/logrus"
	"io"
	"path/filepath"
	"time"
)

type FileObject struct {
	manager        *cloud_disk.Manager
	db             *db.Engine
	storageService *storage.Service
}

func NewFileObject(manager *cloud_disk.Manager) *FileObject {
	storageServ, err := storage.NewService(manager.Config.Storage)
	if err != nil {
		logrus.Fatalf("获取存储服务失败:%s \n", err.Error())
		return nil
	}
	return &FileObject{
		manager:        manager,
		storageService: storageServ,
	}
}

func (s *FileObject) CreateDir(parentID string, dirName string) error {
	// 校验目录名称

	// 先查询parent
	var parent models.FileObject
	err := s.db.DB.Where("id", parentID).First(&parent).Error
	if err != nil {
		return err
	}

	// dirName:=

	object.ID = utils.NewHashID()
	object.CreatedAt = time.Now().UnixMilli()
	object.IsDir = true
	if parent.ObjName != "" {
		object.ObjName = filepath.Join(parent.ObjName, object.ObjName)
	}
	s.storageService.UploadFile()
	return nil
}

func (s *FileObject) CreateFile(file io.Reader) {

}

func (s *FileObject) checkDirName(dirName string) error {

}
