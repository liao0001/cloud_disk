package utils

import (
	"encoding/json"
	"fmt"
	"log"
)

/*接口类型断言*/

//接口对象转为int64
func ObjectToInt64(obj interface{}, def ...int64) int64 {
	d := int64(0)
	if len(def) > 0 {
		d = def[0]
	}

	i, ok := obj.(int64)
	if !ok {
		f, o := obj.(float64)
		if !o {
			return d
		}
		return int64(f)
	}
	return i
}

//接口对象转为int
func ObjectToInt(obj interface{}, def ...int) int {
	d := 0
	if len(def) > 0 {
		d = def[0]
	}
	i, ok := obj.(int)
	if !ok {
		f, o := obj.(float64)
		if !o {
			return d
		}
		return int(f)
		return d
	}
	return i
}

//接口对象转为float64
func ObjectToFloat64(obj interface{}, def ...float64) float64 {
	d := float64(0)
	if len(def) > 0 {
		d = def[0]
	}
	i, ok := obj.(float64)
	if !ok {
		return d
	}
	return i
}

func ObjectToString(obj interface{}, def ...string) string {
	d := ""
	if len(def) > 0 {
		d = def[0]
	}
	i, ok := obj.(string)
	if !ok {
		return d
	}
	return i
}

func ObjectBsToStringArr(obj interface{}, def ...string) []string {
	if obj == nil {
		return []string{}
	}
	ii, ok := obj.([]string)
	if ok {
		return ii
	}
	bi, ok := obj.([]byte)
	if !ok {
		return []string{}
	}

	var res []string
	json.Unmarshal(bi, &res)
	return res
}

func ObjectToStringArr(obj interface{}) []string {
	if obj == nil {
		return []string{}
	}
	ii, ok := obj.([]string)
	if ok {
		return ii
	}
	ib, ok := obj.([]byte)
	if ok {
		var ss []string
		json.Unmarshal(ib, &ss)
		return ss
	}
	ia, ok := obj.([]interface{})
	if ok {
		res := make([]string, 0, len(ia))
		for i := 0; i < len(ia); i++ {
			s, _ := ia[i].(string)
			res = append(res, s)
		}
		return res
	}
	return []string{}
}

func ObjectToMapArr(obj interface{}) (res []map[string]interface{}) {
	if obj == nil {
		return
	}

	var err error
	var bs []byte
	var ok bool
	if bs, ok = obj.([]byte); !ok {
		bs, err = json.Marshal(obj)
		if err != nil {
			return
		}
	}

	err = json.Unmarshal(bs, &res)
	if err != nil {
		log.Printf("反序列化数据数组失败:" + err.Error())
		return
	}
	return
}

func ObjectToMap(obj interface{}) (res map[string]interface{}) {
	if obj == nil {
		return
	}

	if v, o := obj.(map[string]interface{}); o {
		return v
	}

	var err error
	var bs []byte
	var ok bool
	if bs, ok = obj.([]byte); !ok {
		bs, err = json.Marshal(obj)
		if err != nil {
			return
		}
	}

	err = json.Unmarshal(bs, &res)
	if err != nil {
		log.Printf("反序列化数据失败:" + err.Error())
		return
	}
	return
}

//类型转换(强转)
func InterfaceArrToStringArr(ss []interface{}) []string {
	if len(ss) == 0 {
		return []string{}
	}
	res := make([]string, len(ss))
	for k, is := range ss {
		res[k] = fmt.Sprint(is)
	}
	return res
}
