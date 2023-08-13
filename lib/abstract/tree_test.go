package abstract

import (
	"encoding/json"
	"fmt"
	"testing"
)

type DepartmentTest struct {
	ID       string `json:"id"`
	ParentID string `json:"parent_id"`
	Name     string `json:"name"`
	Order    int    `json:"order"`
}

func (d DepartmentTest) GetID() string {
	return d.ID
}

func (d DepartmentTest) GetParentID() string {
	return d.ParentID
}

func TestITreeArr_ToTree(t *testing.T) {
	deptArr := []DepartmentTest{
		{"000", "", "总公司", 0},
		{"010", "000", "分公司1", 0},
		{"020", "000", "分公司2", 0},
		{"011", "010", "1部门", 0},
		{"021", "020", "2部门", 0},
		{"013", "010", "3部门", 0},
	}

	trees := ITreeArr[DepartmentTest](deptArr).ToTree()
	bs, err := json.Marshal(&trees)
	if err != nil {
		fmt.Println("---出错啦:", err)
		//panic(err)
	}
	fmt.Println("结果：", string(bs))
}
