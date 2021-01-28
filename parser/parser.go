package parser

import (
	"github.com/tealeg/xlsx/v3"
	"github.com/toledompm/kroton-etl-go/helper"
)

/*
SheetParseOptions are passed to the parse module to define what will be changed.
headerDictionary: map used to translate file header, keys = old values, values = new values
colParseOptions: map containing parse options for each column. Columns are identified by their header (map keys).
*/
type SheetParseOptions struct {
	headerDictionary map[string]string
	colParseOptions  map[string]int
}

func readFirstSheet(filePath string) (*xlsx.Sheet, error) {
	workBook, err := xlsx.OpenFile(filePath)

	if err != nil {
		return nil, err
	}

	sheet := workBook.Sheets[0]

	return sheet, err
}

func translateHeader(header *xlsx.Row, dict map[string]string) (*xlsx.Row, map[string]int, error) {

	columnIndexes := make(map[string]int)

	header.ForEachCell(
		func(cell *xlsx.Cell) error {
			colName, err := cell.FormattedValue()

			if err != nil {
				return err
			}

			newColName := dict[helper.Normalize(colName)]

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
