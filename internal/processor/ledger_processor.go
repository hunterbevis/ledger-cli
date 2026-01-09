package processor

import (
	"fmt"
	"sort"
	"time"

	"github.com/hunterbevis/ledger-cli/internal/logging"
	"github.com/hunterbevis/ledger-cli/internal/models"
)

type ledgerProcessor struct {
    reporter logging.Reporter
}

func NewProcessor(r logging.Reporter) Processor {
    return &ledgerProcessor{
        reporter: r,
    }
}

func (p *ledgerProcessor) Process(transactions []models.Transaction, period string) (models.Statement, error) {
	
    targetTime, err := time.Parse("200601", period)
    if err != nil {
        return models.Statement{}, logging.ErrInvalidPeriod
    }

    var filtered []models.Transaction
    var income, expenditure int

    for _, t := range transactions {
        if t.Date.Year() == targetTime.Year() && t.Date.Month() == targetTime.Month() {
            filtered = append(filtered, t)

            if t.Amount > 0 {
                income += t.Amount
            } else {
                expenditure += t.Amount
            }
        }
    }

    if len(filtered) == 0 && len(transactions) > 0 {
        p.reporter.ReportProcessWarning(fmt.Sprintf("no transactions found for period %s (checked %d records)", period, len(transactions)))
    }

    sort.Slice(filtered, func(i, j int) bool {
        return filtered[i].Date.After(filtered[j].Date)
    })

    outputPeriod := fmt.Sprintf("%d/%02d", targetTime.Year(), targetTime.Month())

    return models.Statement{
        Period:           outputPeriod,
        TotalIncome:      income,
        TotalExpenditure: expenditure,
        Transactions:     filtered,
    }, nil
}
