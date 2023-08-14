package models

type FileObject struct {
	ID          string `json:"id" gorm:"size:50"`
	StorageTyp  string `json:"storage_typ" gorm:"size:50"` // 存储介质
	IsDir       bool   `json:"is_dir"`                     // 是否为目录
	FileTyp     string `json:"file_typ" gorm:"size:50"`    // 文件类型  暂时不用
	ParentID    string `json:"parent_id" gorm:"size:50"`   // 上级id
	Name        string `json:"name" gorm:"size:255"`       // 文件[夹]名称（显示名称）
	Description string `json:"description" db:"size:500"`  // 文件[夹]描述
	ObjName     string `json:"obj_name" gorm:"size:255"`   // 上传时使用的objName
	Link        string `json:"link" gorm:"size:500"`       // 文件地址
	LinkThumb   string `json:"link_thumb" gorm:"size:500"` // 缩略图
	CreatorID   string `json:"creator_id" gorm:"size:50"`  // 提交人
	CreatedAt   int64  `json:"created_at"`                 // 添加时间
}

//
func (*FileObject) TableName() string {
	return "file_objects"
}
