package parser

import "github.com/tealeg/xlsx/v3"

//TranslateHeader takes the header row and parses each cell value according to a dictionary provided
func TranslateHeader(row *xlsx.Row, dict map[string]string) (*xlsx.Row, error) {
	row.ForEachCell(
		func(cell *xlsx.Cell) error {
			cellText, err := cell.FormattedValue()

			if err != nil {
				return err
			}

			cell.SetString(
				dict[cellText],
			)
			return nil
		},
	)

	return row, nil
}
