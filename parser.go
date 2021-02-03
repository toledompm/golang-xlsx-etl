package parser

import (
	"fmt"

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
	colDictionary map[string]string
	callback      func(*xlsx.Cell) error
}

func translateCell(cell *xlsx.Cell, dict map[string]string) error {
	normalizedCellName := util.Normalize(cell.Value)

	newColName, ok := dict[normalizedCellName]

	if !ok {
		return fmt.Errorf("Key: %s not found in dictionary", normalizedCellName)
	}

	cell.SetString(
		newColName,
	)
	return nil
}

func readFirstSheet(filePath string) (*xlsx.Sheet, error) {
	workBook, err := xlsx.OpenFile(filePath)

	if err != nil {
		return nil, err
	}

	sheet := workBook.Sheets[0]

	return sheet, nil
}

func translateHeader(header *xlsx.Row, dict map[string]string) (map[int]string, error) {
	columnIndexes := make(map[int]string)

	header.ForEachCell(
		func(cell *xlsx.Cell) error {
			err := translateCell(cell, dict)

			if err != nil {
				return err
			}

			col, _ := cell.GetCoordinates()
			columnIndexes[col] = cell.Value
			return nil
		},
	)

	return columnIndexes, nil
}

func parseRow(
	row *xlsx.Row,
	parseOptions map[string]ColumnParseOptions,
	columnIndex map[int]string,
) error {
	row.ForEachCell(func(cell *xlsx.Cell) error {
		colIndex, _ := cell.GetCoordinates()
		colName, ok := columnIndex[colIndex]

		if !ok {
			return fmt.Errorf("No parseOptions for column: %s", colName)
		}

		colParseOptions := parseOptions[colName]

		err := translateCell(cell, colParseOptions.colDictionary)

		if err != nil {
			fmt.Println(err)
		}

		return colParseOptions.callback(cell)
	})

	return nil
}

//Parse parses a xlsx file located at FilePath based on ParseOptions
func Parse(filePath string, parseOptions SheetParseOptions) {
}
