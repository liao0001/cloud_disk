package sexcel

import (
	"encoding/csv"
	"io/ioutil"
	"os"
	"strings"
)

// map格式的数据   标题和数据都是map格式的数组，数据按照 key值映射
// 使用时需使用 New系列方法创建，直接使用结构体的话可能会造成数据错误的问题
type ExcelDataCsv struct {
	Titles [][]interface{}
	Rows   [][]interface{}

	maxIndex int
}

func NewExcelDataCsv(csvPath string, titleLength int) (*ExcelDataCsv, error) {
	fileInfo, err := os.Stat(csvPath)
	if err != nil {
		return nil, err
	}
	//大文件特殊处理
	if fileInfo.Size() > 2<<10 {

	}

	dbs, _ := ioutil.ReadFile(csvPath)
	return NewExcelDataCsvData(string(dbs), titleLength)
}

func NewExcelDataCsvData(csvData string, titleLength int) (*ExcelDataCsv, error) {
	o := &ExcelDataCsv{}
	err := o.initData(csvData, titleLength)
	if err != nil {
		return nil, err
	}
	o.setMaxIndex()
	return o, nil
}

func (o *ExcelDataCsv) initData(csvData string, titleLength int) error {
	records, err := csv.NewReader(strings.NewReader(csvData)).ReadAll()
	if err != nil {
		return err
	}
	if len(records) < titleLength {
		return nil
	}

	o.Titles = o.parseRecords(records[:titleLength])
	o.Rows = o.parseRecords(records[titleLength:])
	return nil
}

func (o *ExcelDataCsv) GetTitles() [][]interface{} {
	return o.Titles
}

func (o *ExcelDataCsv) GetData() [][]interface{} {
	return o.Rows
}

func (o *ExcelDataCsv) GetWidth() int {
	return o.maxIndex + 1
}

//基础校验  有title的时候按照title最大长度显示，没有的时候，按照数据最大列数补全
func (o *ExcelDataCsv) setMaxIndex() {
	if o.maxIndex > 0 {
		return
	}
	if len(o.Titles) > 0 {
		for _, row := range o.Titles {
			if len(row) > o.maxIndex {
				o.maxIndex = len(row)
			}
		}
	}
	if o.maxIndex == 0 {
		for _, row := range o.Rows {
			if len(row) > o.maxIndex {
				o.maxIndex = len(row)
			}
		}
	}
}

//矩阵化数据
func (o *ExcelDataCsv) matrixTitleAndData() {
	o.setMaxIndex()
	//补全标题行
	for i := 0; i < len(o.Titles); i++ {
		if len(o.Titles[i]) < o.maxIndex {
			extend := make([]interface{}, o.maxIndex-len(o.Titles[i]))
			o.Titles[i] = append(o.Titles[i], extend...)
		}
	}

	//补全数据行
	for i := 0; i < len(o.Rows); i++ {
		if len(o.Rows[i]) < o.maxIndex {
			extend := make([]interface{}, o.maxIndex-len(o.Rows[i]))
			o.Rows[i] = append(o.Rows[i], extend...)
		} else if len(o.Rows[i]) > o.maxIndex {
			o.Rows[i] = o.Rows[i][:o.maxIndex]
		}
	}
}

//类型转换
func (o *ExcelDataCsv) parseRecords(lines [][]string) [][]interface{} {
	res := make([][]interface{}, len(lines))
	for key, record := range lines {
		res[key] = o.parseRecord(record)
	}
	return res
}

//类型转换
func (o *ExcelDataCsv) parseRecord(line []string) []interface{} {
	res := make([]interface{}, len(line))
	for key, v := range line {
		res[key] = v
	}
	return res
}
