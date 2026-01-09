package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/hunterbevis/ledger-cli/internal/formatter"
	"github.com/hunterbevis/ledger-cli/internal/parser"
	"github.com/hunterbevis/ledger-cli/internal/processor"
)

func main() {

	period := flag.String("period", "", "The target month in YYYYMM format (e.g., 200601).")
	filePath := flag.String("file", "", "The path to the CSV transaction file.")
	flag.Parse()

	if *period == "" || *filePath == "" {
		fmt.Println("Usage: ledger-cli -period YYYYMM -file path/to/file.csv")
		os.Exit(1)
	}

	p := parser.NewCSVParser()
	transactions, err := p.Parse(*filePath)
	if err != nil {
		fmt.Printf("Error parsing file: %v\n", err)
		os.Exit(1)
	}

	proc := processor.NewProcessor()
	statement, err := proc.Process(transactions, *period)
	if err != nil {
		fmt.Printf("Error processing data: %v\n", err)
		os.Exit(1)
	}

	f := formatter.NewJSONFormatter()
	output, err := f.Format(statement)
	if err != nil {
		fmt.Printf("Error formatting output: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(string(output))
}