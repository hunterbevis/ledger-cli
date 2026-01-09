package models

// JSONTransactionDTO represents the string-based format for the JSON output.
type JSONTransactionDTO struct {
	Date    string `json:"date"`
	Amount  string `json:"amount"` // Per sample: amount is a string.
	Content string `json:"content"`
}

// JSONStatementDTO represents the final JSON structure for the ledger output.
type JSONStatementDTO struct {
	Period           string               `json:"period"`
	TotalIncome      int                  `json:"total_income"`
	TotalExpenditure int                  `json:"total_expenditure"`
	Transactions     []JSONTransactionDTO `json:"transactions"`
}