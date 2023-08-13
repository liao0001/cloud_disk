package abstract

//获取指定key的集合  一般获取id name等
//注: 此接口一个结构生成的是固定的数组 需要支持多字段转数组的，可使用  IKeyValue
type IID interface {
	GeID() string
}

//注: 此接口一个结构生成的是固定的数组 需要支持多字段转数组的，可使用  IKeyArr
type IIDArr[V IID] []V

func (arr IIDArr[V]) Keys() []string {
	res := make([]string, 0, len(arr))
	for _, v := range arr {
		res = append(res, v.GeID())
	}
	return res
}

func (arr IIDArr[V]) KeysUnique() []string {
	res := make([]string, 0, len(arr))
	for _, v := range arr {
		res = appendUnique(res, v.GeID())
	}
	return res
}
