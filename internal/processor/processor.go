package processor

import "github.com/hunterbevis/ledger-cli/internal/models"

// Processor defines the interface for calculating and filtering financial statements.
type Processor interface {
	Process(transactions []models.Transaction, period string) (models.Statement, error)
}
