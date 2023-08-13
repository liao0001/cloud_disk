package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

// 数据库配置
type Config struct {
	Driver      string `yaml:"driver"`
	Url         string `yaml:"url"`
	IdleConn    int    `yaml:"idle_conn"`
	OpenConn    int    `yaml:"open_conn"`
	MaxLifetime int    `yaml:"max_lifetime"`
}

type Engine struct {
	DB      *gorm.DB
	isClone bool
}

func NewEngine(conf *Config) (*Engine, error) {
	dsn := conf.Url
	var db *gorm.DB
	var err error
	switch conf.Driver {
	case "postgres":
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			SkipDefaultTransaction: true,
			CreateBatchSize:        200,
			// Logger:                 logger.Default.LogMode(logger.Info), //调试打印所有日志
			Logger: logger.Default.LogMode(logger.Warn),
		})
	case "sqlite":
		db, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{
			SkipDefaultTransaction: true,
			CreateBatchSize:        200,
			// Logger:                 logger.Default.LogMode(logger.Info), //调试打印所有日志
			Logger: logger.Default.LogMode(logger.Warn),
		})
	case "mysql":
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			SkipDefaultTransaction: true,
			CreateBatchSize:        200,
			// Logger:                 logger.Default.LogMode(logger.Info), //调试打印所有日志
			Logger: logger.Default.LogMode(logger.Warn),
		})
	}

	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(conf.IdleConn)
	sqlDB.SetMaxOpenConns(conf.OpenConn)
	sqlDB.SetConnMaxLifetime(time.Duration(conf.MaxLifetime) * time.Second)

	return &Engine{
		DB: db,
	}, nil
}

// X系列函数  使用简便的链式调用
func (e *Engine) XModel(value interface{}) *Engine {
	return &Engine{
		DB:      e.DB.Model(value),
		isClone: true,
	}
}

func (e *Engine) XWhere(query interface{}, args ...interface{}) *Engine {
	if e.isClone {
		e.DB = e.DB.Where(query, args...)
		return e
	}
	return &Engine{
		DB:      e.DB.Where(query, args...),
		isClone: true,
	}
}

func (e *Engine) XWithPage(page, pageSize int) *Engine {
	if e.isClone {
		e.DB = e.DB.Offset((page - 1) * pageSize).Limit(pageSize)
		return e
	}
	return &Engine{
		DB:      e.DB.Offset((page - 1) * pageSize).Limit(pageSize),
		isClone: true,
	}
}

func (e *Engine) XOrder(sort string) *Engine {
	if len(sort) == 0 {
		return e
	}
	if e.isClone {
		e.DB = e.DB.Order(sort)
		return e
	}
	return &Engine{
		DB:      e.DB.Order(sort),
		isClone: true,
	}
}

func (e *Engine) XCount() (int, error) {
	var count int64
	err := e.DB.Count(&count).Error
	return int(count), err
}

func (e *Engine) XCount64() (int64, error) {
	var count int64
	err := e.DB.Count(&count).Error
	return count, err
}

func (e *Engine) XMapList() ([]map[string]interface{}, error) {
	res := make([]map[string]interface{}, 0, 50)
	err := e.DB.Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (e *Engine) XMapFirst() ([]map[string]interface{}, error) {
	res := make([]map[string]interface{}, 0, 50)
	err := e.DB.First(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}
