package csv

import (
	"dividend-visualizer/models"
	"fmt"
	"strconv"
)

// ParseCSV converts raw CSV records into a DataSet.
func ParseCSV(records [][]string) (models.DataSet, error) {
	var dataSet models.DataSet

	for i, row := range records {
		if i == 0 {
			continue // ヘッダーをスキップ
		}

		// 日付、金額の順に格納されている
		// row: [2024/12/30 71.53] が格納されている
		date := row[0]

		dividend, err := strconv.ParseFloat(row[1], 64)
		if err != nil {
			return nil, fmt.Errorf("invalid dividend format on row %d: %w", i, err)
		}

		data := models.Data{
			DepositDate: date,
			Dividend:    dividend,
		}
		dataSet = append(dataSet, data)
	}

	return dataSet, nil
}
