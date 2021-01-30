package parser

import (
	"github.com/tealeg/xlsx/v3"
	"github.com/toledompm/kroton-etl-go/util"
)

/*
SheetParseOptions are passed to the parse module to define what will be changed.

headerDictionary: map used to translate file header. **Keys will be normalized**

colParseOptions: map containing parse options for each column. Columns are identified by their header (map keys).
*/
type SheetParseOptions struct {
	headerDictionary   map[string]string
	colParseOptionsMap map[string]ColumnParseOptions
}

/*
ColumnParseOptions will be applied to each cell contained inside a given column.

colType: Excel data type to be applied

colDictionary: map used to translate every cell value contained in the column. **Keys will be normalized**

callback: custom function to applied to each cell in a column
*/
type ColumnParseOptions struct {
	colType       string
	colDictionary map[string]string
	callback      func()
}

func readFirstSheet(filePath string) (*xlsx.Sheet, error) {
	workBook, err := xlsx.OpenFile(filePath)

	if err != nil {
		return nil, err
	}

	sheet := workBook.Sheets[0]

	return sheet, nil
}

func translateHeader(header *xlsx.Row, dict map[string]string) (*xlsx.Row, map[string]int, error) {
	columnIndexes := make(map[string]int)

	header.ForEachCell(
		func(cell *xlsx.Cell) error {
			colName, err := cell.FormattedValue()

			if err != nil {
				return err
			}

			newColName := dict[util.Normalize(colName)]

			if newColName == "" {
				newColName = colName
			}

			columnIndexes[newColName], _ = cell.GetCoordinates()

			cell.SetString(
				newColName,
			)
			return nil
		},
	)

	return header, columnIndexes, nil
}

//Parse parses a xlsx file located at FilePath based on ParseOptions
func Parse(filePath string, parseOptions SheetParseOptions) {
}
