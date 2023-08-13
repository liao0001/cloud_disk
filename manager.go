package cloud_disk

import (
	"github.com/liao0001/cloud_disk/config"
	"github.com/liao0001/cloud_disk/db"
	"github.com/liao0001/cloud_disk/storage"
)

type Manager struct {
	DB         *db.Engine
	StorageMap *storage.Factory
	Config     config.Conf
}

type ManagerOption func(manager *Manager)

func NewManager(opts ...ManagerOption) *Manager {
	m := &Manager{}
	for _, opt := range opts {
		opt(m)
	}
	return m
}
