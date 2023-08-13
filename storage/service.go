package storage

import "io"

type Service struct {
	conf    Config
	storage IStorage
}

func NewService(conf Config) (*Service, error) {
	newFunc, err := FactoryGet(conf.DefaultStorage.Key)
	if err != nil {
		return nil, err
	}

	storage, err := newFunc(conf.DefaultStorage)
	if err != nil {
		return nil, err
	}

	return &Service{
		conf:    conf,
		storage: storage,
	}, nil
}

// 上传文件
func (s *Service) UploadFile(objName string, reader io.Reader) (string, error) {
	resLink, err := s.storage.PushFile(objName, reader)
	return resLink, err
}

// 添加文件夹
func (s *Service) CreateDir(objName string) error {
	return s.storage.CreateDir(objName)
}
