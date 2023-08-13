package utils

type Filter struct {
	PageSize     int    `json:"page_size"`
	CurrPage     int    `json:"curr_page"`
	Offset       int    `json:"offset" `
	ParentID     int    `json:"parent_id"`
	DepartmentID int    `json:"department_id"`
	UserID       string `json:"user_id" `
	Type         int    `json:"type" ` // 部门： 0 部门列表 1 角色列表
	IsGroup      bool   `json:"is_group"`
}
