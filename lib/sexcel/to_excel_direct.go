package sexcel

// map格式的数据   标题和数据都是map格式的数组，数据按照 key值映射
// 使用时需使用 New系列方法创建，直接使用结构体的话可能会造成数据错误的问题
type ExcelDataDirect struct {
	Titles [][]interface{}
	Rows   [][]interface{}

	maxIndex int
}

func NewExcelDataDirect(titles [][]interface{}, rows [][]interface{}) *ExcelDataDirect {
	o := &ExcelDataDirect{
		Titles: titles,
		Rows:   rows,
	}
	o.setMaxIndex()
	return o
}

func NewExcelDataDirectOneTitle(title []interface{}, rows [][]interface{}) *ExcelDataDirect {
	o := &ExcelDataDirect{
		Titles: [][]interface{}{title},
		Rows:   rows,
	}
	o.matrixTitleAndData()

	return o
}

func (o *ExcelDataDirect) GetTitles() [][]interface{} {
	return o.Titles
}

func (o *ExcelDataDirect) GetData() [][]interface{} {
	return o.Rows
}

func (o *ExcelDataDirect) GetWidth() int {
	return o.maxIndex + 1
}

//基础校验  有title的时候按照title最大长度显示，没有的时候，按照数据最大列数补全
func (o *ExcelDataDirect) setMaxIndex() {
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
func (o *ExcelDataDirect) matrixTitleAndData() {
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
