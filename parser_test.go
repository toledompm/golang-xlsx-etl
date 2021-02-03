package parser

import (
	"testing"

	"github.com/tealeg/xlsx/v3"
	"github.com/toledompm/kroton-etl-go/util"
)

func TestTranslateHeader(t *testing.T) {
	headerDictionary := mockCellDictionary()

	headerValues := mockCellValues()
	headerRow := mockRow(headerValues)

	util.NormalizeMapKeys(headerDictionary)

	cIndex, _ := translateHeader(headerRow, headerDictionary)

	headerRow.ForEachCell(func(c *xlsx.Cell) error {
		cellText := c.Value
		cellCol, _ := c.GetCoordinates()

		if returnedColName, ok := cIndex[cellCol]; ok {
			if returnedColName != cellText {
				t.Errorf(
					"Wrong column translated, expect column index: %d to return %s, got %s",
					cellCol,
					cellText,
					returnedColName,
				)
			}
		} else if cellText != "UNUSEDCOL" {
			t.Errorf("Column index for %s not found, expected index: %d", cellText, cellCol)
		}

		return nil
	})

}

func TestParseRow(t *testing.T) {
	headerDictionary := mockCellDictionary()
	headerRow := mockRow(mockCellValues())
	util.NormalizeMapKeys(headerDictionary)
	cIndex, _ := translateHeader(headerRow, headerDictionary)

	cParseOpts := make(map[string]ColumnParseOptions)

	callbackCallCount := 0
	callback := func(c *xlsx.Cell) error {
		callbackCallCount++

		_, found := util.Find(mockCellTranslatedValues(), c.Value)

		if !found {
			t.Errorf("Cell with value: %s was not translated", c.Value)
		}
		return nil
	}

	baseColParseOptions := mockColParseOptions()
	baseColParseOptions.callback = callback

	for _, colName := range cIndex {
		cParseOpts[colName] = baseColParseOptions
	}

	row := mockRow(mockCellValues())

	parseRow(row, cParseOpts, cIndex)

	if callbackCallCount != len(cIndex) {
		t.Errorf("Callback was called only: %d, expected: %d", callbackCallCount, len(cIndex))
	}
}

func mockCellDictionary() map[string]string {
	return map[string]string{
		"TEXTCOL":   "text",
		"NumberCOL": "number",
		"dateCol":   "date",
	}
}

func mockCellValues() []string {
	return []string{"TEXTCOL", "NumberCol", "dateCol", "UNUSEDCOL"}
}

func mockCellTranslatedValues() []string {
	return []string{"text", "number", "date"}
}

func mockSheet() *xlsx.Sheet {
	wb := xlsx.NewFile()

	sheet, _ := wb.AddSheet("TestSheet")

	return sheet
}

func mockRow(mockValues []string) *xlsx.Row {
	headerRow := mockSheet().AddRow()

	for counter, key := range mockValues {
		headerRow.GetCell(counter).SetValue(key)
	}

	return headerRow
}

func mockColParseOptions() ColumnParseOptions {
	return ColumnParseOptions{
		colDictionary: util.NormalizeMapKeys(mockCellDictionary()),
	}
}
