package util

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
	"reflect"
	"time"
)

const (
	A = iota + 'A'
)

type StructExcel struct {
	f       *excelize.File
	Sheet1  string
	saveto  string
	winline int
}

func NewDefaultStructExcel() StructExcel {
	return StructExcel{
		f:       excelize.NewFile(),
		Sheet1:  "Sheet1",
		saveto:  "excel.xlsx",
		winline: 1,
	}
}

// SetHeaderByStruct 设置表头
// 传入结构体对象,和想要设置的行,把结构体表头映射到excel上
// structobj 传入结构体实例
// begin 开始的位置
func (e *StructExcel) SetHeaderByStruct(structobj interface{}, begin int) {
	if begin == 0 {
		begin = 1
	}
	var getValue = reflect.TypeOf(structobj)

	for i := 0; i < getValue.NumField(); i++ {
		structTitle := getValue.Field(i).Name
		_ = e.f.SetCellValue(e.Sheet1, fmt.Sprintf("%c", A+i)+fmt.Sprintf("%d",
			begin), structTitle)
	}

}

// SetHeaderInfo 设置excel头部备注
func (e *StructExcel) SetHeaderInfo(info string, x, y string) {
	styleId, _ := e.f.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{
			Horizontal: "center", //水平居中
			Vertical:   "center", //垂直居中
		},
	})
	e.f.MergeCell(e.Sheet1, x, y)
	e.f.SetCellStyle(e.Sheet1, x, y, styleId)
	e.f.SetCellValue(e.Sheet1, x, info)
}

// RangeStructSlice 序列化struct为excel
// structSlice参数一：要序列化的结构体切片
// begin参数二：开始的位置
func (e *StructExcel) RangeStructSlice(structSlice interface{}, begin int) {

	style, _ := e.f.NewStyle(`{"number_format": 0}`)

	getValue := reflect.ValueOf(structSlice)
	if getValue.Kind() != reflect.Slice {
		panic("need slice kind")
	}
	fmt.Println()
	fmt.Println()
	l := getValue.Len()
	ll := 0
	for i := 0; i < l; i++ {
		value := getValue.Index(i)
		typel := value.Type()

		if typel.Kind() != reflect.Struct {
			panic("need struct kind")
		}
		num := value.NumField()
		ll = num
		for j := 0; j < num; j++ {
			fmt.Printf("name: %s, type: %s, value: %v\n", typel.Field(j).Name, value.Field(j).Type(), value.Field(j).Interface())
			if lstr := Filter(value.Field(j)); lstr == "" {
				e.f.SetCellValue(e.Sheet1, fmt.Sprintf("%c", A+j)+fmt.Sprintf("%d", begin+i), value.Field(j).Interface())
			} else {
				e.f.SetCellValue(e.Sheet1, fmt.Sprintf("%c", A+j)+fmt.Sprintf("%d", begin+i), lstr)

			}
		}

	}
	e.f.SetCellStyle(e.Sheet1, fmt.Sprintf("%c", A)+fmt.Sprintf("%d", begin),
		fmt.Sprintf("%c", A+l)+fmt.Sprintf("%d", begin+ll), style)

}

// ExcelSaveAs 保存excel的地址
func (e *StructExcel) ExcelSaveAs(path string) {
	if path != "" {
		e.saveto = path
	}

	// 根据指定路径保存文件
	if err := e.f.SaveAs(e.saveto); err != nil {
		fmt.Println(err)
	}
}

// ExportDefaultExcel 序列化结构体 到 excel
//参数一是结构体
//参数二是要序列化成excel的结构体切片
//参数三是保存路径
func ExportDefaultExcel(structobj interface{}, structSlice interface{}, path string) {
	excel := StructExcel{
		f:       excelize.NewFile(),
		Sheet1:  "Sheet1",
		saveto:  "excel.xlsx",
		winline: 1,
	}

	excel.SetHeaderByStruct(structobj, 1)
	excel.RangeStructSlice(structSlice, 2)
	excel.ExcelSaveAs(path)
}

// Filter 过滤器
func Filter(value reflect.Value) string {
	switch value.Type() {
	case reflect.TypeOf(time.Time{}):
		t := value.Interface().(time.Time)
		return t.Format("2006.01.02 15:04:05")
		//return "time11"
	case reflect.TypeOf(gorm.DeletedAt{}):
		return "-"
	}
	return ""
}
