package db

import (
	"fmt"
	"testing"
)

func TestSqliteJson(t *testing.T) {
	engineCase, err := NewEngine(&Config{
		Driver:      "sqlite",
		Url:         `data.db`,
		IdleConn:    10,
		OpenConn:    10,
		MaxLifetime: 60,
	})
	if err != nil {
		panic(err)
	}
	sql := fmt.Sprintf(`select * from files`)
	var res []string
	err = engineCase.DB.Raw(sql).Find(&res).Error
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
