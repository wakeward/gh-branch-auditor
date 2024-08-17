package output

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	"github.com/wakeward/gh-branch-auditor/pkg/eval"
)

// JSONWriter implements result Writer
type JSONWriter struct {
	Output io.Writer
}

func PrettyJSON(jsonBytes []byte) ([]byte, error) {
	var out bytes.Buffer
	err := json.Indent(&out, jsonBytes, "", "  ")
	if err != nil {
		return jsonBytes, err
	}
	return out.Bytes(), nil
}

// Write writes the reports in JSON format
func (jw JSONWriter) Write(reports eval.Reports) error {
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
