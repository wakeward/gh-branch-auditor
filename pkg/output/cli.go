package output

import (
	"fmt"
	"io"

	"github.com/wakeward/gh-branch-auditor/pkg/eval"
)

// CLIOutput implements result Writer
type CLIOutput struct {
	Output io.Writer
}

// Write writes the reports in CLI format
func (jw CLIOutput) Write(reports eval.Reports) error {

	fmt.Println("\nBranch Auditor Results:")
	// tm := "%-25s %-10s %-30s\n"

	// for _, p := range reports {
	// 	fmt.Printf("\nRepository: %s\n\n", p.Repo)
	// 	fmt.Printf(tm, "ID", "Risk", "Reason")
	// 	for _, r := range p.Rules {
	// 		fmt.Printf(tm, r.ID, r.Risk, r.Reason)
	// 	}
	// }
	return nil
}
