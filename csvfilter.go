// Script to take CSV input and output only specific columns
// usage:
// cat data | csvfilter -c 1,4

package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var (
	version = flag.String("version", "", "print the version number")
	lazy = flag.Bool("lazy", false, "allow lazy csv formats (variable number of columns)")
	columnsStr = flag.String("c", "", "Index (0 based) for which CSV columns to display (comma separate multiple)")
)

func main() {
	flag.Parse()
	
	var columns []int
	
	for _, s := range strings.Split(*columnsStr, ",") {
		i, err := strconv.Atoi(s)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Invalid argument for -c %v", s)
			os.Exit(1)
		}
		columns = append(columns, i)
	}

	if len(columns) == 0 {
		fmt.Println("missing arg -c")
		os.Exit(1)
	}

	reader := csv.NewReader(os.Stdin)
	if *lazy {
		reader.FieldsPerRecord = -1
	}
	writer := csv.NewWriter(os.Stdout)
	for {
		var out []string
		row, err := reader.Read()
		if len(row) > 0 {
			for _, index := range columns {
				if len(row) <= index {
					fmt.Fprintf(os.Stderr, "Error: row has %d columns, expected %d\n", len(row), index)
					writer.Flush()
					os.Exit(1)
				}
				out = append(out, row[index])
			}
			writeErr := writer.Write(out)
			if writeErr != nil {
				fmt.Fprintf(os.Stderr, "Error: %s\n", err)
				os.Exit(2)
			}
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			writer.Flush()
			fmt.Fprintf(os.Stderr, "Error: %s\n", err)
			os.Exit(3)
		}
	}
	writer.Flush()
	err := writer.Error()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(4)
	}

}
