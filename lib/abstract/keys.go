package abstract

//获取指定key的集合  一般获取id name等
type IKey interface {
	GetValueByKey(keyName string) string
}

//注: 此接口一个结构生成的是固定的数组 需要支持多字段转数组的，可使用  IKeyArr
type IKeyArr[V IKey] []V

func (arr IKeyArr[V]) Keys(keyName string) []string {
	res := make([]string, 0, len(arr))
	for _, v := range arr {
		res = append(res, v.GetValueByKey(keyName))
	}
	return res
}

func (arr IKeyArr[V]) KeysUnique(keyName string) []string {
	res := make([]string, 0, len(arr))
	for _, v := range arr {
		res = appendUnique(res, v.GetValueByKey(keyName))
	}
	return res
}
