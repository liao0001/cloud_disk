package abstract

type IIDValue interface {
	GetID() string
	GetName() string
}

//id和value
type IDValue struct {
	ID   string `json:"id"`
	Name string `json:"value"`
}

type IIDValueArr[V IIDValue] []V

//数组直接转为 IDValue 数组
func (arr IIDValueArr[V]) ToIDValueArr() IDValueArr {
	res := make(IDValueArr, 0, len(arr))
	for i := 0; i < len(arr); i++ {
		res = append(res, IDValue{
			ID:   arr[i].GetID(),
			Name: arr[i].GetName(),
		})
	}
	return res
}

type IDValueArr []IDValue

//带额外数据的id value
type IDValueOther struct {
	ID    string         `json:"id"`
	Value string         `json:"value"`
	Other map[string]any `json:"other"`
}
