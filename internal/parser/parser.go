package parser

import "github.com/hunterbevis/ledger-cli/internal/models"

// Parser defines the interface for validating and parsing through input files (csv, json, etc...).
type Parser interface {
	Parse(path string) ([]models.Transaction, error)
}
