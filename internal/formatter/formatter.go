package formatter

import "github.com/hunterbevis/ledger-cli/internal/models"

// Formatter defines the interface for presenting statements in specific formats.
type Formatter interface {
	Format(statement models.Statement) ([]byte, error)
}
