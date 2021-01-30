package parser

import (
	"fmt"
	"testing"

	"github.com/tealeg/xlsx/v3"
)

func testSheet() *xlsx.Sheet {
	wb := xlsx.NewFile()

	sheet, _ := wb.AddSheet("TestSheet")

	return sheet
}

func TestTranslateHeader(t *testing.T) {
	headerDictionary := map[string]string{
		"textCol":   "text",
		"numberCol": "number",
		"dateCol":   "date",
	}

	headerRow := testSheet().AddRow()

	fmt.Println(headerRow, headerDictionary)
}
