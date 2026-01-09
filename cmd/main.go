package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/hunterbevis/ledger-cli/internal/formatter"
	"github.com/hunterbevis/ledger-cli/internal/logging"
	"github.com/hunterbevis/ledger-cli/internal/parser"
	"github.com/hunterbevis/ledger-cli/internal/processor"
)

func main() {
	period := flag.String("period", "", "the target month in YYYYMM format (e.g., 200601).")
	filePath := flag.String("file", "", "the path to the transaction file.")
	flag.Parse()

	if *period == "" || *filePath == "" {
		fmt.Println("Usage: ledger-cli -period YYYYMM -file path/to/file.csv")
		os.Exit(1)
	}

	rep := logging.NewStandardReporter()

	p := parser.NewCSVParser(rep)

	transactions, err := p.Parse(*filePath)
	if err != nil {
		os.Exit(1)
	}

	proc := processor.NewProcessor(rep)
	statement, err := proc.Process(transactions, *period)
	if err != nil {
		os.Exit(1)
	}

	f := formatter.NewJSONFormatter(rep)
	output, err := f.Format(statement)
	if err != nil {
		os.Exit(1)
	}

	fmt.Println(string(output))
}
