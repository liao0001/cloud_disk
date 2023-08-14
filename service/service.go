package service

import "github.com/liao0001/cloud_disk"

type Service struct {
	manager *cloud_disk.Manager
}

func NewService(manager *cloud_disk.Manager) *Service {
	return &Service{manager: manager}
}

func (s *Service) FileObject() *FileObjectService {
	return NewFileObjectService(s.manager)
}
