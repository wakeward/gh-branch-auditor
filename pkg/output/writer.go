package output

import (
	"io"

	"github.com/wakeward/gh-branch-auditor/pkg/eval"
)

// type reports eval.Reports

// WriteReports writes the result to output, format as passed in argument
func WriteReports(format string, output io.Writer, reports eval.Reports) error {
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
	Write(eval.Reports) error
}
