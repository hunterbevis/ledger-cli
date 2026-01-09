package formatter

import (
	"encoding/json"
	"fmt"

	"github.com/hunterbevis/ledger-cli/internal/models"
)

type jsonFormatter struct{}

func NewJSONFormatter() Formatter {
	return &jsonFormatter{}
}

func (f *jsonFormatter) Format(statement models.Statement) ([]byte, error) {
	output := models.JSONStatementDTO{
		Period:           statement.Period,
		TotalIncome:      statement.TotalIncome,
		TotalExpenditure: statement.TotalExpenditure,
		Transactions:     make([]models.JSONTransactionDTO, len(statement.Transactions)),
	}

	for i, t := range statement.Transactions {
		output.Transactions[i] = models.JSONTransactionDTO{
			Date:    t.Date.Format("2006/01/02"),
			Amount:  fmt.Sprintf("%d", t.Amount),
			Content: t.Content,
		}
	}

	return json.MarshalIndent(output, "", "  ")
}
