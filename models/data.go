package models

// Data represents a single row of CSV data.
type Data struct {
	DepositDate string  // 日付
	Dividend    float64 // 金額
}

// DataSet represents multiple rows of CSV data.
type DataSet []Data
