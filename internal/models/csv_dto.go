package models

// Constants for expected CSV column names.
const (
	CSVColumnDate    = "date"
	CSVColumnAmount  = "amount"
	CSVColumnContent = "content"
)

// CSVHeaderDTO represents the structure of the CSV header for validation.
type CSVHeaderDTO struct {
	Columns []string
}

// GetExpectedColumns returns the required column names in the correct order.
func (h CSVHeaderDTO) GetExpectedColumns() []string {
	return []string{CSVColumnDate, CSVColumnAmount, CSVColumnContent}
}

// CSVTransactionDTO matches the raw record structure from the CSV file.
type CSVTransactionDTO struct {
	Date    string
	Amount  string
	Content string
}
