package sexcel

import (
	"fmt"
	"github.com/xuri/excelize/v2"
)

const (
	MergeTypeNone MergeType = iota //不合并
	MergeTypeH                     //行合并
	MergeTypeV                     //列合并
	MergeTypeHV                    //行+列合并
)

//合并类型
type MergeType int

//单元格合并的设置
type CellMergeSetting struct {
	Typ        MergeType     //合并类型
	RangeAll   bool          //不限范围
	RangeIndex []int         //范围  -1为不限 行合并时，这里是行下标  列合并时，这里是列下标  行和列都有时(先行后列): all = [-1,-1,-1,-1]
	Rule       CellMergeRule //合并规则 默认是 DefaultMergeRule
}

//合并规则
type CellMergeRule func(a, b interface{}) bool

//默认 nil和 "" 相等
var DefaultMergeRule CellMergeRule = func(a, b interface{}) bool {
	var as string
	if a == nil {
		as = ""
	} else {
		as = fmt.Sprintf("%v", a)
	}
	var bs string
	if b == nil {
		bs = ""
	} else {
		bs = fmt.Sprintf("%v", b)
	}
	return as == bs
}

//进行合并
func (m *CellMergeSetting) MergeCells(file *excelize.File, sheetName string) error {
	//使用默认的判断方式
	if m.Rule == nil {
		m.Rule = DefaultMergeRule
	}

	var err error
	switch m.Typ {
	case MergeTypeH:
		err = m.mergeCellsH(file, sheetName)
		if err != nil {
			return err
		}
	case MergeTypeV:
		err = m.mergeCellsV(file, sheetName)
		if err != nil {
			return err
		}
	case MergeTypeHV:
		err = m.mergeCellsH(file, sheetName)
		if err != nil {
			return err
		}
		err = m.mergeCellsV(file, sheetName)
		if err != nil {
			return err
		}
	}
	return nil
}

//水平方向合并
func (m *CellMergeSetting) mergeCellsH(file *excelize.File, sheetName string) error {
	var start int
	var end int
	if m.RangeAll {
		start, end = -1, -1
	} else if len(m.RangeIndex) == 2 {
		start, end = m.RangeIndex[0], m.RangeIndex[1]
	} else {
		return nil
	}

	rows, err := file.Rows(sheetName)
	if err != nil {
		return fmt.Errorf("获取rows失败:%s ", err.Error())
	}
	//遍历每一行
	for rowIdx := 0; rows.Next() && (rowIdx <= end || end == -1); rowIdx++ {
		//判断是否在设定范围内
		if rowIdx < start {
			continue
		}

		//进行行合并
		cells, err := rows.Columns()
		if err != nil {
			return err
		}

		if len(cells) < 2 {
			continue
		}

		slow := 0
		for fast := range cells {
			//这里使用合并规则合并
			if !m.Rule(cells[fast], cells[slow]) {
				//出现不同值的时候，触发修改，fast-1之前都是相同的
				if fast-slow > 1 {
					err = file.MergeCell(sheetName, GetAxis(slow, rowIdx), GetAxis(fast-1, rowIdx))
					if err != nil {
						return err
					}
				}
				slow = fast
			}
			if fast == len(cells)-1 {
				//结束的时候也要判断下
				if fast-slow > 0 {
					err = file.MergeCell(sheetName, GetAxis(slow, rowIdx), GetAxis(fast, rowIdx))
					if err != nil {
						return err
					}
				}
			}
		}
	}
	return nil
}

//垂直方向合并
func (m *CellMergeSetting) mergeCellsV(file *excelize.File, sheetName string) error {
	var start int
	var end int
	if m.RangeAll {
		start, end = -1, -1
	} else if len(m.RangeIndex) == 2 {
		start, end = m.RangeIndex[0], m.RangeIndex[1]
	} else {
		return nil
	}

	cols, err := file.Cols(sheetName)
	if err != nil {
		return fmt.Errorf("获取rows失败:%s ", err.Error())
	}

	//遍历每一行
	for colIdx := 0; cols.Next() && (colIdx <= end || end == -1); colIdx++ {
		//判断是否在设定范围内
		if colIdx < start {
			continue
		}

		//进行行合并
		cells, err := cols.Rows()
		if err != nil {
			return err
		}

		if len(cells) < 2 {
			continue
		}

		slow := 0
		for fast := range cells {
			//这里使用合并规则合并
			if !m.Rule(cells[fast], cells[slow]) {
				//出现不同值的时候，触发修改，fast-1之前都是相同的
				if fast-slow > 1 {
					err = file.MergeCell(sheetName, GetAxis(colIdx, slow), GetAxis(colIdx, fast-1))
					if err != nil {
						return err
					}
				}
				slow = fast
			}
			if fast == len(cells)-1 {
				//结束的时候也要判断下
				if fast-slow > 0 {
					err = file.MergeCell(sheetName, GetAxis(colIdx, slow), GetAxis(colIdx, fast))
					if err != nil {
						return err
					}
				}
			}
		}
	}
	return nil
}
