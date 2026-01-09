package parser

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/hunterbevis/ledger-cli/internal/models"
)

type csvParser struct{}

func NewCSVParser() Parser {
	return &csvParser{}
}

func (p *csvParser) Parse(path string) ([]models.Transaction, error) {
	
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("Failed to open file: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	rawHeader, err := reader.Read()
	if err != nil {
		return nil, fmt.Errorf("Failed to read CSV header: %w", err)
	}

	headerDTO := models.CSVHeaderDTO{Columns: rawHeader}
	if err := validateHeader(headerDTO); err != nil {
		return nil, err
	}

	var transactions []models.Transaction

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("Error reading CSV row: %w", err)
		}

		if err := validateRecord(record); err != nil {
			return nil, err
		}

		dto := models.CSVTransactionDTO{
			Date:    strings.TrimSpace(record[0]),
			Amount:  strings.TrimSpace(record[1]),
			Content: strings.TrimSpace(record[2]),
		}

		t, err := mapDTOToModel(dto)
		if err != nil {
			return nil, err
		}

		transactions = append(transactions, t)
	}

	return transactions, nil
}

func mapDTOToModel(dto models.CSVTransactionDTO) (models.Transaction, error) {
	parsedDate, err := time.Parse("2006/01/02", dto.Date)
	if err != nil {
		return models.Transaction{}, fmt.Errorf("Invalid date format %s: %w", dto.Date, err)
	}

	amount, err := strconv.Atoi(dto.Amount)
	if err != nil {
		return models.Transaction{}, fmt.Errorf("Invalid amount %s: %w", dto.Amount, err)
	}

	return models.Transaction{
		Date:    parsedDate,
		Amount:  amount,
		Content: dto.Content,
	}, nil
}

func validateHeader(dto models.CSVHeaderDTO) error {

	expected := dto.GetExpectedColumns()

	if len(dto.Columns) != len(expected) {
		return fmt.Errorf("Invalid CSV schema: expected %d columns, got %d", len(expected), len(dto.Columns))
	}

	for i, name := range expected {
		actual := strings.ToLower(strings.TrimSpace(dto.Columns[i]))
		if actual != name {
			return fmt.Errorf("Invalid header: expected column %d to be '%s' but got '%s'", i+1, name, dto.Columns[i])
		}
	}
	return nil
}

func validateRecord(record []string) error {

	if len(record) < 3 {
		return fmt.Errorf("Invalid row found: insufficient columns")
	}

	for i, value := range record {
		if strings.TrimSpace(value) == "" {
			return fmt.Errorf("Invalid row found: column %d is empty", i+1)
		}
	}
	return nil
}
