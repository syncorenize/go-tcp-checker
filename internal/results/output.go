package results

import (
	"encoding/json"
	"fmt"
	"os"
)

// OutputFormat specifies the desired format of the output.
type OutputFormat string

const (
	FormatText OutputFormat = "text"
	FormatJSON OutputFormat = "json"
)

// PrintResults prints or writes the service check results in the specified format.
func PrintResults(results []ServiceResult, format OutputFormat, outputFile string) {
	var output string

	switch format {
	case FormatJSON:
		jsonOutput, err := json.MarshalIndent(results, "", "  ")
		if err != nil {
			fmt.Println("Error formatting JSON output:", err)
			return
		}
		output = string(jsonOutput)
	default: // Default to text format
		for _, result := range results {
			status := "available"
			if !result.Available {
				status = "unavailable"
				if result.Error != "" {
					status += fmt.Sprintf(" (%s)", result.Error)
				}
			}
			output += fmt.Sprintf("Service at %s is %s\n", result.Address, status)
		}
	}

	// Decide whether to print to console or write to file based on outputFile
	if outputFile == "" {
		fmt.Print(output)
	} else {
		fmt.Println("Output file:", outputFile)
		err := os.WriteFile(outputFile, []byte(output), 0644)
		if err != nil {
			fmt.Println("Error writing to file:", err)
		}
	}
}
