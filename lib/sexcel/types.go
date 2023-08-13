package sexcel

//
type ExcelDataTitleObject struct {
	Key   string `json:"key"`
	Title string `json:"title"`
}

type ExcelDataTitleObjectArr []ExcelDataTitleObject

func (arr ExcelDataTitleObjectArr) ToMapTitle() (map[string]interface{}, map[string]int) {
	m := make(map[string]interface{}, len(arr))
	idx := make(map[string]int, len(arr))
	for k, v := range arr {
		m[v.Key] = v.Title
		idx[v.Key] = k
	}
	return m, idx
}
