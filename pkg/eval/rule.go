package eval

import (
	"github.com/wakeward/gh-branch-auditor/pkg/branchprotections"
)

type Rule struct {
	ID        string
	Rule      string
	Risk      string
	Severity  string
	Predicate func(*branchprotections.RepoBranchProtection) int
}

// Eval executes the predicate if the kind matches the rule
func (r *Rule) Eval(bpr *branchprotections.RepoBranchProtection) int {
	count := r.Predicate(bpr)
	return count
}
