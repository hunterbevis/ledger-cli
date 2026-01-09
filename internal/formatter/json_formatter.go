package formatter

import (
	"encoding/json"
	"fmt"

	"github.com/hunterbevis/ledger-cli/internal/logging"
	"github.com/hunterbevis/ledger-cli/internal/models"
)

type jsonFormatter struct {
    reporter logging.Reporter
}

func NewJSONFormatter(r logging.Reporter) Formatter {
    return &jsonFormatter{
        reporter: r,
    }
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

    data, err := json.MarshalIndent(output, "", "  ")
    if err != nil {
        f.reporter.ReportFormatError(logging.ErrFormatFailed)
        return nil, logging.ErrFormatFailed
    }

    return data, nil
}
