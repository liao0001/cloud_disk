package sexcel

//表格数据
// 数据行中直接使用 time.Time的时候需要使用UTC时区的
type IToExcelData interface {
	GetTitles() [][]interface{} //标题行
	GetData() [][]interface{}   //数据行
	GetWidth() int              //矩阵的数据的宽
}

//装饰器
type ExcelResDecoration func(res *ExcelParserResToExcel) *ExcelParserResToExcel
