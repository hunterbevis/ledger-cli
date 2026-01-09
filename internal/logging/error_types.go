package logging

import (
	"errors"
	"fmt"
)

var (
    ErrFileNotFound     = errors.New("the specified file could not be found")
    ErrInvalidHeader    = errors.New("the file header does not match the expected schema")
    ErrEmptyFile        = errors.New("the file contains no data")
    ErrInsufficientCols = errors.New("row has fewer columns than expected")
    ErrEmptyColumn      = errors.New("a required column is empty")
    ErrInvalidDate      = errors.New("date is not in the required YYYY/MM/DD format")
    ErrInvalidAmount    = errors.New("amount must be a valid integer")
    ErrInvalidPeriod    = errors.New("the period must be in yyyymm format (e.g., 202601)")
    ErrFormatFailed     = errors.New("failed to format the statement output")
)

type FileError struct {
	Path string
	Err  error
}

func (e *FileError) Error() string {
	return fmt.Sprintf("file error: %v (path: %s)", e.Err, e.Path)
}

func (e *FileError) Unwrap() error {
	return e.Err
}

type RowError struct {
	Line  int
	Err   error
	Value string
}

func (e *RowError) Error() string {
	return fmt.Sprintf("line %d: %v", e.Line, e.Err)
}

func (e *RowError) Unwrap() error {
	return e.Err
}
