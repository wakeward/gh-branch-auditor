package eval

import "github.com/wakeward/gh-branch-auditor/pkg/branchprotections"

type Rule struct {
	ID        string
	Reason    string
	Risk      string
	Predicate func([]*branchprotections.RepoBranchProtection) []*branchprotections.RepoBranchProtection
}

// Eval executes the predicate if the kind matches the rule
func (r *Rule) Eval(bpr []*branchprotections.RepoBranchProtection) []*branchprotections.RepoBranchProtection {
	protectionRules := r.Predicate(bpr)
	return protectionRules
}
