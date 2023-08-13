package local

import "github.com/liao0001/cloud_disk/storage"

func init() {
	storage.RegisterFunc("local", NewStorage)
}
