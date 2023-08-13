package models

type FileObject struct {
	ID         string `json:"id"`
	StorageTyp string `json:"storage_typ"` // 存储介质
	IsDir      bool   `json:"is_dir"`      // 是否为目录
	FileTyp    string `json:"file_typ"`    // 文件类型  暂时不用
	ParentID   string `json:"parent_id"`   // 上级id
	Name       string `json:"name"`        // 文件名称（显示名称）
	ObjName    string `json:"obj_name"`    // 上传时使用的objName
	Link       string `json:"link"`        // 文件地址
	LinkThumb  string `json:"link_thumb"`  // 缩略图
	CreatorID  string `json:"creator_id"`  //
	CreatedAt  int64  `json:"created_at"`  // 添加时间
}

//
func (*FileObject) TableName() string {
	return "file_objects"
}
