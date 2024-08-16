package output

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/wakeward/gh-branch-auditor/pkg/eval"
)

var Now = time.Now

type reports eval.Reports

// WriteReports writes the result to output, format as passed in argument
func WriteReports(format string, output io.Writer, reports reports) error {
	var writer Writer
	switch format {
	case "json":
		writer = &JSONWriter{Output: output}
	case "cli":
		writer = &CLIOutput{Output: output}
	}

	if err := writer.Write(reports); err != nil {
		return err
	}
	return nil
}

// Writer defines the result write operation
type Writer interface {
	Write(reports) error
}

// JSONWriter implements result Writer
type JSONWriter struct {
	Output io.Writer
}

// CLIOutput implements result Writer
type CLIOutput struct {
	Output io.Writer
}

// PrettyJSON will indent JSON to be pretty
func PrettyJSON(jsonBytes []byte) ([]byte, error) {
	var out bytes.Buffer
	err := json.Indent(&out, jsonBytes, "", "  ")
	if err != nil {
		return jsonBytes, err
	}
	return out.Bytes(), nil
}

// Write writes the reports in JSON format
func (jw JSONWriter) Write(reports reports) error {
	output, err := json.Marshal(reports)
	if err != nil {
		return err
	}

	formattedOutput, err := PrettyJSON(output)
	if err != nil {
		return err
	}
	if _, err = fmt.Fprint(jw.Output, string(formattedOutput)); err != nil {
		return err
	}
	return nil
}

// Write writes the reports in CLI format
func (jw CLIOutput) Write(reports reports) error {

	fmt.Println("\nBranch Auditor Results:")
	tm := "%-40s %-10s %-10s\n"

	for _, p := range reports {
		for _, report := range p.Rules {
			fmt.Printf("\nRisk: %s\nReason: %s\n", report.Risk, report.Reason)
			fmt.Printf(tm, "Repository", "Branch", "Protected")
			for _, pr := range report.ProtectionRules {

				fmt.Printf(tm, pr.RepoName, pr.Branch, strconv.FormatBool(pr.IsProtected))
			}
		}

	}
	return nil
}
