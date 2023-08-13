package sexcel

import (
	"encoding/json"
	"fmt"
	"github.com/xuri/excelize/v2"
	"io/ioutil"
	"testing"
	"time"
)

func TestToExcel(t *testing.T) {
	title := []interface{}{"姓名", "班级", "年龄", "出生日期"}
	rows := [][]interface{}{
		{"张三", "3年1班", 12, time.Date(2010, 01, 02, 0, 0, 0, 0, time.Local)},
		{"李四", nil, 11, time.Date(2011, 01, 02, 0, 0, 0, 0, time.Local)},
		{"王五", "3年3班", nil, time.Date(2012, 01, 02, 0, 0, 0, 0, time.Local)},
		{"赵六", "3年4班", 9, nil},
		{"qiqi", "3年4班", 10, time.Date(2012, 01, 02, 0, 0, 0, 0, time.Local), "呵呵呵"},
	}

	data := NewExcelDataDirectOneTitle(title, rows)

	err := NewExcelParser("测试数据表").ToExcel(data).
		WriteToFile(`D:\workspace\dev\src\git.siruijie.com.cn\pkg\sexcel\test\基础表.xlsx`)
	if err != nil {
		panic(err)
	}
}

func TestToExcelMerge(t *testing.T) {
	data, err := NewExcelDataCsv(`D:\workspace\dev\src\git.siruijie.com.cn\pkg\sexcel\test\test.csv`, 2)
	if err != nil {
		panic(err)
	}

	src := NewExcelParser("测试数据表2").ToExcel(data)

	err = src.WriteToFile(`D:\workspace\dev\src\git.siruijie.com.cn\pkg\sexcel\test\合并文件-基础.xlsx`)

	err = src.AddTitleMerge().AddMergeSetting(CellMergeSetting{
		Typ:        MergeTypeV,
		RangeAll:   false,
		RangeIndex: []int{0, 1},
		Rule:       DefaultMergeRule,
	}).WriteToFile(`D:\workspace\dev\src\git.siruijie.com.cn\pkg\sexcel\test\合并文件.xlsx`)
	if err != nil {
		panic(err)
	}
}

func TestStyleSet(t *testing.T) {
	_ = excelize.Style{
		Border: []excelize.Border{ //边框-数组-上右下左
			{
				Type:  "", //边框类型
				Color: "", //颜色
				Style: 0,  //样式id
			},
		},
		Fill: excelize.Fill{ //单元格填充
			Type:    "",         //填充类型
			Pattern: 0,          //
			Color:   []string{}, //颜色
			Shading: 0,
		},
		Font: &excelize.Font{ //字体
			Bold:      false,
			Italic:    false,
			Underline: "",
			Family:    "",
			Size:      0,
			Strike:    false,
			Color:     "",
		},
		Alignment: &excelize.Alignment{ //对齐方式
			Horizontal:      "", //水平方向
			Indent:          0,  //缩进
			JustifyLastLine: false,
			ReadingOrder:    0,
			RelativeIndent:  0,
			ShrinkToFit:     false,
			TextRotation:    0,
			Vertical:        "",
			WrapText:        false,
		},
		Protection:    nil,
		NumFmt:        0, //数字格式
		DecimalPlaces: 0, //小数位
		CustomNumFmt:  nil,
		Lang:          "",
		NegRed:        false,
	}
	//file := excelize.NewFile()
	//file.SetConditionalFormat()
	//cases := []struct {
	//	label      string
	//	format     string
	//	expectFill bool
	//}{{
	//	label:      "no_fill",
	//	format:     `{"alignment":{"wrap_text":true}}`,
	//	expectFill: false,
	//}, {
	//	label:      "fill",
	//	format:     `{"fill":{"type":"pattern","pattern":1,"color":["#000000"]}}`,
	//	expectFill: true,
	//}}
	//
	//for _, testCase := range cases {
	//	xl := excelize.NewFile()
	//	styleID, err := xl.NewStyle(testCase.format)
	//	assert.NoError(t, err)
	//
	//	styles := xl.stylesReader()
	//	style := styles.CellXfs.Xf[styleID]
	//	if testCase.expectFill {
	//		assert.NotEqual(t, *style.FillID, 0, testCase.label)
	//	} else {
	//		assert.Equal(t, *style.FillID, 0, testCase.label)
	//	}
	//}
	//f := NewFile()
	//styleID1, err := f.NewStyle(`{"fill":{"type":"pattern","pattern":1,"color":["#000000"]}}`)
	//assert.NoError(t, err)
	//styleID2, err := f.NewStyle(`{"fill":{"type":"pattern","pattern":1,"color":["#000000"]}}`)
	//assert.NoError(t, err)
	//assert.Equal(t, styleID1, styleID2)
	//assert.NoError(t, f.SaveAs(filepath.Join("test", "TestStyleFill.xlsx")))
}

func TestToExcelFromMap(t *testing.T) {
	var mapData struct {
		Titles      []map[string]interface{} `json:"titles"`
		TitleIndexs []map[string]int         `json:"title_indexs"`
		Rows        []map[string]interface{} `json:"rows"`
	}
	bs, _ := ioutil.ReadFile(`D:\workspace\dev\src\git.siruijie.com.cn\data\bi\app\lib\sexcel\test.json`)
	_ = json.Unmarshal(bs, &mapData)

	eserv, err := NewExcelDataMap(mapData.Titles, mapData.Rows, mapData.TitleIndexs)
	if err != nil {
		panic(err)
	}
	err = NewExcelParser("数据表").ToExcel(eserv).WriteToFile(`D:\workspace\dev\src\git.siruijie.com.cn\data\bi\app\lib\sexcel\test.xlsx`)
	fmt.Println(err)
}

//测试数据
var testData = `[
      {
        "f_1606286132624": "东华",
        "f_1606286144445": "一年级",
        "f_1606286167432": 54.25,
        "f_1606286175675": 8
      },
      {
        "f_1606286132624": "东华",
        "f_1606286144445": "三年级",
        "f_1606286167432": 71.5,
        "f_1606286175675": 10
      },
      {
        "f_1606286132624": "东华",
        "f_1606286144445": "二年级",
        "f_1606286167432": 92,
        "f_1606286175675": 9
      },
      {
        "f_1606286132624": "东华",
        "f_1606286144445": "五年级",
        "f_1606286167432": null,
        "f_1606286175675": null
      },
      {
        "f_1606286132624": "东华",
        "f_1606286144445": "四年级",
        "f_1606286167432": null,
        "f_1606286175675": null
      },
      {
        "f_1606286132624": "中通",
        "f_1606286144445": "一年级",
        "f_1606286167432": 78,
        "f_1606286175675": 9
      },
      {
        "f_1606286132624": "中通",
        "f_1606286144445": "三年级",
        "f_1606286167432": 83,
        "f_1606286175675": 10
      },
      {
        "f_1606286132624": "中通",
        "f_1606286144445": "二年级",
        "f_1606286167432": null,
        "f_1606286175675": null
      },
      {
        "f_1606286132624": "中通",
        "f_1606286144445": "五年级",
        "f_1606286167432": 54,
        "f_1606286175675": 8
      },
      {
        "f_1606286132624": "中通",
        "f_1606286144445": "四年级",
        "f_1606286167432": 78.2,
        "f_1606286175675": 10
      },
      {
        "f_1606286132624": "北环",
        "f_1606286144445": "一年级",
        "f_1606286167432": 86.25,
        "f_1606286175675": 10
      },
      {
        "f_1606286132624": "北环",
        "f_1606286144445": "三年级",
        "f_1606286167432": 65.5,
        "f_1606286175675": 10
      },
      {
        "f_1606286132624": "北环",
        "f_1606286144445": "二年级",
        "f_1606286167432": 75.11111111111111,
        "f_1606286175675": 11
      },
      {
        "f_1606286132624": "北环",
        "f_1606286144445": "五年级",
        "f_1606286167432": 62.5,
        "f_1606286175675": 7
      },
      {
        "f_1606286132624": "北环",
        "f_1606286144445": "四年级",
        "f_1606286167432": 79,
        "f_1606286175675": 11
      },
      {
        "f_1606286132624": "南屿",
        "f_1606286144445": "一年级",
        "f_1606286167432": 75.6,
        "f_1606286175675": 11
      },
      {
        "f_1606286132624": "南屿",
        "f_1606286144445": "三年级",
        "f_1606286167432": 73.85714285714286,
        "f_1606286175675": 10
      },
      {
        "f_1606286132624": "南屿",
        "f_1606286144445": "二年级",
        "f_1606286167432": 77.375,
        "f_1606286175675": 10
      },
      {
        "f_1606286132624": "南屿",
        "f_1606286144445": "五年级",
        "f_1606286167432": 83.83333333333333,
        "f_1606286175675": 11
      },
      {
        "f_1606286132624": "南屿",
        "f_1606286144445": "四年级",
        "f_1606286167432": 85.66666666666667,
        "f_1606286175675": 10
      },
      {
        "f_1606286132624": "西部",
        "f_1606286144445": "一年级",
        "f_1606286167432": 63,
        "f_1606286175675": 7
      },
      {
        "f_1606286132624": "西部",
        "f_1606286144445": "三年级",
        "f_1606286167432": 78.83333333333333,
        "f_1606286175675": 10
      },
      {
        "f_1606286132624": "西部",
        "f_1606286144445": "二年级",
        "f_1606286167432": 72.75,
        "f_1606286175675": 10
      },
      {
        "f_1606286132624": "西部",
        "f_1606286144445": "五年级",
        "f_1606286167432": null,
        "f_1606286175675": null
      },
      {
        "f_1606286132624": "西部",
        "f_1606286144445": "四年级",
        "f_1606286167432": 73.75,
        "f_1606286175675": 11
      },
      {
        "f_1606286132624": "东华",
        "f_1606286144445": "汇总",
        "f_1606286167432": 66.11111111111111,
        "f_1606286175675": 10
      },
      {
        "f_1606286132624": "中通",
        "f_1606286144445": "汇总",
        "f_1606286167432": 77.2,
        "f_1606286175675": 10
      },
      {
        "f_1606286132624": "北环",
        "f_1606286144445": "汇总",
        "f_1606286167432": 75.25,
        "f_1606286175675": 11
      },
      {
        "f_1606286132624": "南屿",
        "f_1606286144445": "汇总",
        "f_1606286167432": 79.09375,
        "f_1606286175675": 11
      },
      {
        "f_1606286132624": "西部",
        "f_1606286144445": "汇总",
        "f_1606286167432": 74.26086956521739,
        "f_1606286175675": 11
      }`
