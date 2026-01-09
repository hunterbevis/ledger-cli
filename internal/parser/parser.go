package parser

import "github.com/hunterbevis/ledger-cli/internal/models"

// Global to tell our reader objects to not enforce row rules. We will handle that with custom reporter
const relaxedFields = -1

// Parser defines the interface for validating and parsing through input files (csv, json, etc...).
type Parser interface {
	Parse(path string) ([]models.Transaction, error)
}
