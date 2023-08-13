### 封装了excel相关操作的一些方法

支持操作:

* 导出
* 支持多种数据源: go 对象、csv、txt、json等
* 支持多中导出方式: 写入文件、写入 writer、快捷写入 httpWriter
* 支持标题行的横向和纵向单元格合并(相同内容的合并) [仅合并非空文本 数字不合并]
* 支持数据行的纵向单元格合并(相同内容的合并) [仅合并非空文本 数字不合并]
* 解析
* 支持直接从 httpRequest中读取文件
* 支持读取合并了单元格的数据(仅支持识别 标题行:横向合并 数据行:纵向合并 且需要指定标题行的宽度,非标题行视为数据行)

依赖:
github.com/xuri/excelize/v2

### 使用
```go
//先创建一个数据对象  这里支持多种方式，以map为例
data := NewExcelDataDirectOneTitle(title, rows)
//或按照常见的 map arr 创建(这里需要确定title对应关系和title的顺序)
err := NewExcelParser("测试数据表").ToExcel(data).WriteToFile("./test.xlsx")
```

### 样式设置
```go
//可以使用对象格式  或者直接使用对应的json结构字符串
_ = &excelize.Style{
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
        Horizontal:      "",
        Indent:          0,
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

例: {"font":{"bold":true},"alignment":{"horizontal":"center"}}
```
