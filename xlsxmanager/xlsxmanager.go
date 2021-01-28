package xlsxmanager

import (
	"github.com/tealeg/xlsx/v3"
)

//ReadFirstSheet opens a xlsx file and returns the sheet with index 0
func ReadFirstSheet(filePath string) (*xlsx.Sheet, error) {
	workBook, err := xlsx.OpenFile(filePath)

	if err != nil {
		return nil, err
	}

	sheet := workBook.Sheets[0]

	return sheet, err
}
