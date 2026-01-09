package models

import "time"

type Transaction struct {
	Date    time.Time
	Amount  int
	Content string
}

type Statement struct {
	Period           string
	TotalIncome      int
	TotalExpenditure int
	Transactions     []Transaction
}
