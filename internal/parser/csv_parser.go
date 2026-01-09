package parser

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/hunterbevis/ledger-cli/internal/logging"
	"github.com/hunterbevis/ledger-cli/internal/models"
)

var _ Parser = (*csvParser)(nil)

type csvParser struct {
	reporter        logging.Reporter
	fieldsPerRecord int
}

func NewCSVParser(r logging.Reporter) Parser {
	return &csvParser{
		reporter:        r,
		fieldsPerRecord: relaxedFields,
	}
}

func (p *csvParser) Parse(path string) ([]models.Transaction, error) {
	file, err := os.Open(path)
	if err != nil {
		fe := &logging.FileError{
			Path: path,
			Err:  fmt.Errorf("%w: %s", logging.ErrFileNotFound, path),
		}
		p.reporter.ReportFileError(fe)
		return nil, fe
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = p.fieldsPerRecord
	lineCount := 1

	rawHeader, err := reader.Read()
	if err != nil {
		if err == io.EOF {
			fe := &logging.FileError{Path: path, Err: logging.ErrEmptyFile}
			p.reporter.ReportFileError(fe)
			return nil, fe
		}
		
		fe := &logging.FileError{Path: path, Err: err}
		p.reporter.ReportFileError(fe)
		return nil, fe
	}

	if err := validateHeader(models.CSVHeaderDTO{Columns: rawHeader}); err != nil {
		fe := &logging.FileError{
			Path: path,
			Err:  fmt.Errorf("%w: %v", logging.ErrInvalidHeader, err),
		}
		p.reporter.ReportFileError(fe)
		return nil, fe
	}

	var transactions []models.Transaction

	for {
		lineCount++
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			re := &logging.RowError{Line: lineCount, Err: err}
			p.reporter.ReportRowError(re)
			return nil, re
		}

		if p.isEmpty(record) {
			continue
		}

		if err := p.validateRow(lineCount, record); err != nil {
			if re, ok := err.(*logging.RowError); ok {
				p.reporter.ReportRowError(re)
			}
			return nil, err
		}

		dto := models.CSVTransactionDTO{
			Date:    strings.TrimSpace(record[0]),
			Amount:  strings.TrimSpace(record[1]),
			Content: strings.TrimSpace(record[2]),
		}

		t, err := mapDTOToModel(dto)
		if err != nil {
			re := &logging.RowError{
				Line:  lineCount,
				Err:   err,
				Value: strings.Join(record, "|"),
			}
			p.reporter.ReportRowError(re)
			return nil, re
		}

		transactions = append(transactions, t)
	}

	return transactions, nil
}

func (p *csvParser) isEmpty(record []string) bool {
	if len(record) == 0 {
		return true
	}
	if len(record) == 1 && strings.TrimSpace(record[0]) == "" {
		return true
	}
	return false
}

func (p *csvParser) validateRow(line int, record []string) error {
	if len(record) < 3 {
		return &logging.RowError{
			Line:  line,
			Err:   logging.ErrInsufficientCols,
			Value: strings.Join(record, ","),
		}
	}

	for i, val := range record {
		if strings.TrimSpace(val) == "" {
			return &logging.RowError{
				Line:  line,
				Err:   logging.ErrEmptyColumn,
				Value: fmt.Sprintf("Col %d", i+1),
			}
		}
	}
	return nil
}

func validateHeader(dto models.CSVHeaderDTO) error {
	expected := dto.GetExpectedColumns()
	if len(dto.Columns) != len(expected) {
		return fmt.Errorf("expected %d columns, got %d", len(expected), len(dto.Columns))
	}

	for i, name := range expected {
		actual := strings.ToLower(strings.TrimSpace(dto.Columns[i]))
		if actual != name {
			return fmt.Errorf("expected column %d to be '%s' but got '%s'", i+1, name, dto.Columns[i])
		}
	}
	return nil
}

func mapDTOToModel(dto models.CSVTransactionDTO) (models.Transaction, error) {
	parsedDate, err := time.Parse("2006/01/02", dto.Date)
	if err != nil {
		return models.Transaction{}, fmt.Errorf("%w: %v", logging.ErrInvalidDate, err)
	}

	amount, err := strconv.Atoi(dto.Amount)
	if err != nil {
		return models.Transaction{}, fmt.Errorf("%w: %v", logging.ErrInvalidAmount, err)
	}

	return models.Transaction{
		Date:    parsedDate,
		Amount:  amount,
		Content: dto.Content,
	}, nil
}
