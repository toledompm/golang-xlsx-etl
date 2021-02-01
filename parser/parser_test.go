package parser

import (
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
		"textcol":   "text",
		"numbercol": "number",
		"datecol":   "date",
	}
	denormalizedHeaderKeys := []string{"TEXTCOL", "NumberCol", "dateCol"}

	expectedColIndex := map[string]int{"text": 0, "number": 1, "date": 2}

	headerRow := testSheet().AddRow()

	for counter, key := range denormalizedHeaderKeys {
		headerRow.GetCell(counter).SetValue(key)
	}

	cIndex, _ := translateHeader(headerRow, headerDictionary)

	headerRow.ForEachCell(func(c *xlsx.Cell) error {
		cellText, error := c.FormattedValue()
		cellCol, _ := c.GetCoordinates()

		if returnedColIndex, ok := cIndex[cellText]; ok {
			if returnedColIndex != cellCol {
				t.Error("translateHeader returned wrong index")
			}
			if returnedColIndex != expectedColIndex[cellText] {
				t.Errorf(
					"translateHeader translated wrong column, expected index to translate: %d, got: %d",
					expectedColIndex[cellText],
					returnedColIndex,
				)
			}
		} else {
			t.Errorf("translateHeder didnt return index for col: %d, value: %s", cellCol, cellText)
		}

		return error
	})

}
