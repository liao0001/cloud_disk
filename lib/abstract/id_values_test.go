package abstract

import (
	"fmt"
	"testing"
)

type foo struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (f foo) GetID() string {
	return f.ID
}

func (f foo) GetName() string {
	return f.Name
}

func TestParseIDValueArr(t *testing.T) {
	var sourceData = []foo{
		{"111", "张三", 18},
		{"222", "李四", 19},
		{"333", "王五", 20},
		{"444", "赵六", 21},
	}

	fmt.Println(IIDValueArr[foo](sourceData).ToIDValueArr())

	//IIDValueArr[Foo](sourceData).ToIDValueArr()
	//
	//sourceData.ToIDValueArr()
}
