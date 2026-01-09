package logging

import (
	"errors"
	"log"
	"os"
)

type Reporter interface {
	ReportRowError(err *RowError)
	ReportFileError(err *FileError)
	ReportProcessWarning(message string)
	ReportFormatError(err error)
}

type StandardReporter struct {
	logger *log.Logger
}

func NewStandardReporter() Reporter {
	return &StandardReporter{
		logger: log.New(os.Stderr, "[LEDGER-LOGGER] ", log.LstdFlags),
	}
}

func (r *StandardReporter) ReportRowError(re *RowError) {
	switch {
	case errors.Is(re.Err, ErrInvalidAmount):
		r.logger.Printf("AMOUNT ERROR | Line %d: check the formatting of %q", re.Line, re.Value)
	case errors.Is(re.Err, ErrInvalidDate):
		r.logger.Printf("DATE ERROR | Line %d: incompatible date found: %q", re.Line, re.Value)
	default:
		r.logger.Printf("ROW ERROR | %v", re)
	}
}

func (r *StandardReporter) ReportFileError(fe *FileError) {
	r.logger.Printf("FATAL FILE ERROR | path: %s | reason: %v", fe.Path, fe.Err)
}

func (r *StandardReporter) ReportProcessWarning(message string) {
    r.logger.Printf("PROCESSOR WARNING | %s", message)
}

func (r *StandardReporter) ReportFormatError(err error) {
    r.logger.Printf("FORMATTER ERROR | failed to generate output: %v", err)
}
