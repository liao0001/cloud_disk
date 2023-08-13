package sexcel

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"io"
)

const (
	DefaultSheetName = "Sheet1"

	DefaultStyleTitle = `{"font":{"bold":true},"alignment":{"horizontal":"center","vertical":"center"}}`
	DefaultStyleCell  = `{"alignment":{"vertical":"center"}}`
)

// 解释器(单sheet的)
type ExcelParser struct {
	data      IToExcelData
	sheetName string
	err       error

	file        *excelize.File
	totalRowNum int // 处理的行 默认按行添加的话，直接使用游标添加 初始值是0
	titleRowNum int // 标题行数量
	dataRowNum  int // 数据行数量
}

// 创建一个sheet的默认 解析器
func NewExcelParser(sheetName string) *ExcelParser {
	c := &ExcelParser{
		sheetName: sheetName,
	}
	return c
}

// 转换文件内容
func (c *ExcelParser) ToExcel(data IToExcelData) *ExcelParserResToExcel {
	var err error
	c.data = data
	if c.data == nil || c.data == IToExcelData(nil) {
		c.err = fmt.Errorf("ToExcel参数数据为空 ")
		return newExcelParserToExcelRes(c)
	}

	c.initExcelFile()
	// 添加标题
	err = c.writeTitles()
	if err != nil {
		c.err = err
		return newExcelParserToExcelRes(c)
	}

	// 添加数据
	err = c.writeData()
	if err != nil {
		c.err = err
		return newExcelParserToExcelRes(c)
	}
	return newExcelParserToExcelRes(c)
}

// 创建基础文件
func (c *ExcelParser) initExcelFile() {
	// 删除默认的 Sheet1
	c.file = excelize.NewFile()
	if c.sheetName != DefaultSheetName {
		c.file.NewSheet(c.sheetName)
		c.file.DeleteSheet(DefaultSheetName)
	}

	// c.totalRowNum = 1  GetAxis+1返回 所以这里不需要从1开始了
}

// 添加标题
func (c *ExcelParser) writeTitles() error {
	var err error
	for _, row := range c.data.GetTitles() {
		for colIdx, val := range row {
			err = c.file.SetCellValue(c.sheetName, GetAxis(colIdx, c.totalRowNum), val)
			if err != nil {
				return err
			}
		}
		c.titleRowNum++
		c.totalRowNum++
	}

	return nil
}

// 添加数据
func (c *ExcelParser) writeData() error {
	var err error
	for _, row := range c.data.GetData() {
		for colIdx, val := range row {
			err = c.file.SetCellValue(c.sheetName, GetAxis(colIdx, c.totalRowNum), val)
			if err != nil {
				return err
			}
		}
		c.dataRowNum++
		c.totalRowNum++
	}
	return nil
}

// 导出的返回值
type ExcelParserResToExcel struct {
	p   *ExcelParser
	err error

	// 是否不使用标题样式  默认使用标题样式
	noTitleStyle bool
	// 标题样式和数据单元格样式(如果没有 titleStyle，则标题使用 cellStyle，如果 cellStyle 也不存在，则都是默认的)
	titleStyleID  int
	cellStyleID   int
	mergeSettings []CellMergeSetting
}

func newExcelParserToExcelRes(c *ExcelParser) *ExcelParserResToExcel {
	// 这里创建默认样式
	styleID, _ := c.file.NewStyle(DefaultStyleTitle)
	cellID, _ := c.file.NewStyle(DefaultStyleCell)

	res := &ExcelParserResToExcel{
		p:            c,
		err:          c.err,
		titleStyleID: styleID,
		cellStyleID:  cellID,
	}
	return res
}

// 第三方库不支持按row设置  这里自己遍历
func (res *ExcelParserResToExcel) SetTitleNoStyle() *ExcelParserResToExcel {
	if res.err != nil || res.p.file == nil {
		return res
	}
	res.noTitleStyle = true
	return res
}

// 设置样式 不设置的时候，使用默认样式
// 默认样式: title: 居中、加粗 rows: 无
func (res *ExcelParserResToExcel) SetStyle(titleStyle, cellStyle string) *ExcelParserResToExcel {
	if res.err != nil || res.p.file == nil {
		return res
	}
	if len(titleStyle) > 0 {
		res.titleStyleID, res.err = res.p.file.NewStyle(titleStyle)
	}
	if len(cellStyle) > 0 {
		res.cellStyleID, res.err = res.p.file.NewStyle(cellStyle)
	}
	return res
}

// 设置样式  这里部分标题和数据  设置的是全部的
func (res *ExcelParserResToExcel) SetStyleAll(style string) *ExcelParserResToExcel {
	if res.err != nil || res.p.file == nil {
		return res
	}
	if len(style) > 0 {
		var styleID int
		styleID, res.err = res.p.file.NewStyle(style)

		res.titleStyleID = styleID
		res.cellStyleID = styleID
	}
	return res
}

// 设置应用单元格合并规则的位置  一般创建标题合并规则(NewTitleMerge())
func (res *ExcelParserResToExcel) SetMerge(mergeSettings []CellMergeSetting) *ExcelParserResToExcel {
	res.mergeSettings = mergeSettings
	return res
}

// 设置默认的行单元格合并
func (res *ExcelParserResToExcel) AddTitleMerge() *ExcelParserResToExcel {
	cm := CellMergeSetting{
		Typ:        MergeTypeH,
		RangeIndex: []int{0, res.p.titleRowNum - 1},
		Rule:       DefaultMergeRule,
	}
	res.AddMergeSetting(cm)
	return res
}

// 添加合并规则
func (res *ExcelParserResToExcel) AddMergeSetting(mergeSetting CellMergeSetting) *ExcelParserResToExcel {
	if res.mergeSettings == nil {
		res.mergeSettings = []CellMergeSetting{mergeSetting}
	} else {
		res.mergeSettings = append(res.mergeSettings, mergeSetting)
	}
	return res
}

// 数据导出
func (res *ExcelParserResToExcel) Write(writer io.Writer) error {
	if res.err != nil {
		return res.err
	}
	if err := res.writeBefore(); err != nil {
		return err
	}
	return res.p.file.Write(writer)
}

// 导出为文件
func (res *ExcelParserResToExcel) WriteToFile(fileName string) error {
	if res.err != nil {
		return res.err
	}
	if err := res.writeBefore(); err != nil {
		return err
	}

	return res.p.file.SaveAs(fileName)
}

// 导出前 根据设置进行相关操作
func (res *ExcelParserResToExcel) writeBefore() error {
	// 样式
	err := res.setStyle()
	if err != nil {
		return err
	}

	// 合并
	err = res.setDataMerge()
	if err != nil {
		return err
	}
	return nil
}

// 设置样式
func (res *ExcelParserResToExcel) setStyle() error {
	err := res.setCellStyle()
	if err != nil {
		return err
	}

	// 如果表单都设置了，这里title会覆盖cell
	err = res.setTitleStyle()
	if err != nil {
		return err
	}
	return nil
}

// 设置标题样式
func (res *ExcelParserResToExcel) setTitleStyle() error {
	if res.noTitleStyle || res.titleStyleID == 0 || res.p.titleRowNum == 0 {
		return nil
	}
	width := res.p.data.GetWidth()
	if width == 0 {
		return nil
	}

	startCell := "A1"
	endCell := GetAxis(width-1, res.p.titleRowNum-1)
	err := res.p.file.SetCellStyle(res.p.sheetName, startCell, endCell, res.titleStyleID)
	if err != nil {
		return fmt.Errorf("设置标题样式失败: %s ", err.Error())
	}
	return nil
}

// 设置数据样式
func (res *ExcelParserResToExcel) setCellStyle() error {
	if res.cellStyleID == 0 || res.p.dataRowNum == 0 {
		return nil
	}
	width := res.p.data.GetWidth()
	if width == 0 {
		return nil
	}

	// 这里设置全表
	startCell := "A1"
	endCell := GetAxis(width-1, res.p.totalRowNum-1)
	err := res.p.file.SetCellStyle(res.p.sheetName, startCell, endCell, res.cellStyleID)
	if err != nil {
		return fmt.Errorf("设置单元格样式失败: %s ", err.Error())
	}
	return nil
}

// 标题的行和列合并 参数为: 行的 start end
func (res *ExcelParserResToExcel) setDataMerge() error {
	var err error
	for i := 0; i < len(res.mergeSettings); i++ {
		err = res.mergeSettings[i].MergeCells(res.p.file, res.p.sheetName)
		if err != nil {
			return fmt.Errorf("合并单元格时发生错误: setting:%v err:%s ", res.mergeSettings[i], err.Error())
		}
	}
	return nil
}
