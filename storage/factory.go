package storage

import (
	"errors"
	"github.com/sirupsen/logrus"
)

var ErrNoService = errors.New("未注册服务")
var ErrInvalidKey = errors.New("未实现的driver")

type Factory struct {
	defaultDriver string
	newFuncMap    map[string]NewStorageFunc
}

func FactoryGet(driver string) (NewStorageFunc, error) {
	if len(instanceFactory.newFuncMap) == 0 {
		logrus.Fatalf("未注册任何存储服务\n")
		return nil, ErrNoService
	}
	serv, ok := instanceFactory.newFuncMap[driver]
	if !ok || serv == nil {
		logrus.Fatalf("当前存储引擎未注册:%s\n", driver)
		return nil, ErrInvalidKey
	}
	return serv, nil
}

func FactoryGetDefault() NewStorageFunc {
	return instanceFactory.newFuncMap[instanceFactory.defaultDriver]
}

var instanceFactory *Factory

func init() {
	instanceFactory = &Factory{
		newFuncMap: map[string]NewStorageFunc{},
	}
}

func RegisterFunc(driver string, f NewStorageFunc) {
	if len(instanceFactory.defaultDriver) == 0 {
		instanceFactory.defaultDriver = driver
	}

	instanceFactory.newFuncMap[driver] = f
}
