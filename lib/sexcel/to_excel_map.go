package sexcel

import (
	"fmt"
	"log"
)

// map格式的数据   标题和数据都是map格式的数组，数据按照 key值映射
// 使用时需使用 New系列方法创建，直接使用结构体的话可能会造成数据错误的问题
// 使用map格式 就以title为主了  map格式数据不支持没有标题的数据
type ExcelDataMap struct {
	titles      []map[string]interface{} //多行的标题  []key:label  label是显示值，key作为标识
	titlesIndex []map[string]int         //每个key的顺序  从0开始   对应的 key:index
	rows        []map[string]interface{} //对应的 []key:value

	excelTitles [][]interface{} //
	excelRows   [][]interface{}

	//多行标题的时候仅按照最后一行标题计算
	maxIndex    int
	titleKeyRow []string
}

//按照最原始的数据新建
func NewExcelDataMap(titles, rows []map[string]interface{}, titleIndex []map[string]int) (*ExcelDataMap, error) {
	ed := ExcelDataMap{
		titles:      titles,
		titlesIndex: titleIndex,
		rows:        rows,
	}

	ed.setMaxIndex()
	return &ed, nil
}

//titles rows 为数据  均是 key:value  orderTitles 为固定的title key值
func NewExcelDataMapExtraOrder(titles, rows []map[string]interface{}, orderTitles [][]string) (*ExcelDataMap, error) {
	titleIndex := make([]map[string]int, 0, len(orderTitles))
	for _, title := range orderTitles {
		ot := make(map[string]int, len(title))
		for idx, v := range title {
			ot[v] = idx
		}
		titleIndex = append(titleIndex, ot)
	}
	ed := ExcelDataMap{
		titles:      titles,
		titlesIndex: titleIndex,
		rows:        rows,
	}
	ed.setMaxIndex()
	return &ed, nil
}

//rows 为数据  orderTitles 为指定的标题，rows是按照最后一行的 标题为key值
func NewExcelDataMapFixedTitles(orderTitles [][]interface{}, rows []map[string]interface{}) (*ExcelDataMap, error) {
	titles := make([]map[string]interface{}, 0, len(orderTitles))
	titleIndex := make([]map[string]int, 0, len(orderTitles))
	for _, title := range orderTitles {
		tt := make(map[string]interface{}, len(title))
		ot := make(map[string]int, len(title))
		for idx, v := range title {
			tt[fmt.Sprint(v)] = v
			ot[fmt.Sprint(v)] = idx
		}
		titles = append(titles, tt)
		titleIndex = append(titleIndex, ot)
	}

	ed := ExcelDataMap{
		titles:      titles,
		titlesIndex: titleIndex,
		rows:        rows,

		excelTitles: orderTitles,
	}
	ed.setMaxIndex()
	return &ed, nil
}

//rows 为数据  orderTitles 为指定的标题，rows是按照最后一行的 标题为key值
func NewExcelDataMapOneTitle(title map[string]interface{}, rows []map[string]interface{}, titleIndex map[string]int) (*ExcelDataMap, error) {
	ed := ExcelDataMap{
		titles:      []map[string]interface{}{title},
		titlesIndex: []map[string]int{titleIndex},
		rows:        rows,
	}

	ed.setMaxIndex()
	return &ed, nil
}

//按照最原始的数据新建
func NewExcelDataMapNoOrder(titles, rows []map[string]interface{}) (*ExcelDataMap, error) {
	titleIndex := make([]map[string]int, 0, len(titles))
	for _, title := range titles {
		ot := make(map[string]int, len(title))
		idx := 0
		for key := range title {
			ot[key] = idx
			idx++
		}
		titleIndex = append(titleIndex, ot)
	}
	ed := ExcelDataMap{
		titles:      titles,
		titlesIndex: titleIndex,
		rows:        rows,
	}
	ed.setMaxIndex()
	return &ed, nil
}

//使用object方式创建  适用于确定标题，对象数据的情况
func NewExcelDataMapByTitles(titles ExcelDataTitleObjectArr, rows []map[string]interface{}) *ExcelDataMap {
	ts, tIdxs := titles.ToMapTitle()

	ed := ExcelDataMap{
		titles:      []map[string]interface{}{ts},
		titlesIndex: []map[string]int{tIdxs},
		rows:        rows,
	}
	ed.setMaxIndex()
	return &ed
}

func (o *ExcelDataMap) GetTitles() [][]interface{} {
	if len(o.titles) != len(o.titlesIndex) {
		log.Printf("[ERROR] titles和对应的index数量不匹配;titles=%d indexs=%d", len(o.titles), len(o.titlesIndex))
		return [][]interface{}{}
	}

	if len(o.excelTitles) > 0 {
		return o.excelTitles
	}

	res := make([][]interface{}, 0, len(o.titles))
	for i := 0; i < len(o.titlesIndex); i++ {
		//这里就算是空行也要加上
		res = append(res, o.parseTitle(o.titlesIndex[i], o.titles[i]))
	}

	o.setMaxIndex()
	o.excelTitles = res
	return res
}

//
func (o *ExcelDataMap) parseTitle(idxMap map[string]int, titleMap map[string]interface{}) []interface{} {
	if len(idxMap) == 0 || len(titleMap) == 0 {
		log.Printf("[WARNING] 当前标题行和标题顺序行都为空")
		return nil
	}

	row := make([]interface{}, o.maxIndex+1)
	for key, v := range titleMap {
		ii, ok := idxMap[key]
		if ok {
			row[ii] = v
		}
	}
	return row
}

//获取并设置标题的最大下标 这里也是数据的最大下标  超过这个界限的不管了
func (o *ExcelDataMap) setMaxIndex() {
	for _, row := range o.titlesIndex {
		for _, idx := range row {
			if idx > o.maxIndex {
				o.maxIndex = idx
			}
		}
	}
}

//直接按照最后一行取标题的key
func (o *ExcelDataMap) setTitleRow() {
	if len(o.titleKeyRow) > 0 {
		return
	}
	o.titleKeyRow = make([]string, o.maxIndex+1)
	//直接取最后一行标题作为key值映射的标题
	trIndex := len(o.titles) - 1
	for key := range o.titles[trIndex] {
		if ii, ok := o.titlesIndex[trIndex][key]; ok {
			o.titleKeyRow[ii] = key
		}
	}
}

func (o *ExcelDataMap) GetData() [][]interface{} {
	if len(o.excelRows) > 0 {
		return o.excelRows
	}
	o.setTitleRow()

	rows := make([][]interface{}, 0, len(o.rows))
	for _, one := range o.rows {
		row := make([]interface{}, 0, o.maxIndex)
		for _, key := range o.titleKeyRow {
			if len(key) > 0 { //为空的不显示
				row = append(row, one[key])
			}
		}
		rows = append(rows, row)
	}
	o.excelRows = rows

	return rows
}

func (o *ExcelDataMap) GetWidth() int {
	o.setMaxIndex()

	return o.maxIndex + 1
}
