package service

import (
	"github.com/liao0001/cloud_disk"
	"github.com/liao0001/cloud_disk/db"
	"github.com/liao0001/cloud_disk/lib/utils"
	"github.com/liao0001/cloud_disk/models"
	"github.com/liao0001/cloud_disk/storage"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"io"
	"time"
)

type FileObjectService struct {
	manager        *cloud_disk.Manager
	db             *db.Engine
	storageService *storage.Service
}

func NewFileObjectService(manager *cloud_disk.Manager) *FileObjectService {
	storageServ, err := storage.NewService(manager.Config.Storage)
	if err != nil {
		logrus.Fatalf("获取存储服务失败:%s \n", err.Error())
		return nil
	}
	return &FileObjectService{
		manager:        manager,
		db:             manager.DB,
		storageService: storageServ,
	}
}

func (s *FileObjectService) CreateDir(userId, parentID string, dirName string) error {
	// 校验目录名称

	// 先查询parent
	var parent models.FileObject
	err := s.db.DB.Where("id", parentID).First(&parent).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}

	objName := dirName
	if len(parent.ObjName) > 0 {
		objName = parent.ObjName + "/" + objName
	} else {
		objName = userId + "/" + objName
	}

	err = s.storageService.CreateDir(objName)
	if err != nil {
		return errors.Wrap(err, "创建目录失败")
	}

	fileObject := &models.FileObject{
		ID:          utils.NewHashID(),
		StorageTyp:  s.storageService.GetStorageTyp(),
		IsDir:       true,
		FileTyp:     "",
		ParentID:    parentID,
		Name:        dirName,
		Description: "",
		ObjName:     objName,
		Link:        "",
		LinkThumb:   "",
		CreatorID:   userId,
		CreatedAt:   time.Now().UnixMilli(),
	}
	err = s.db.DB.Create(&fileObject).Error
	if err != nil {
		return errors.Wrap(err, "创建目录数据失败")
	}
	return nil
}

func (s *FileObjectService) CreateFile(file io.Reader) {

}

func (s *FileObjectService) checkDirName(dirName string) error {
	return nil
}
