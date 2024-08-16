package rules

import (
	"github.com/wakeward/gh-branch-auditor/pkg/branchprotections"
)

func IsProtected(input []*branchprotections.RepoBranchProtection) []*branchprotections.RepoBranchProtection {
	protectionRule := make([]*branchprotections.RepoBranchProtection, 0)

	for _, rule := range input {
		if !rule.IsProtected {
			protectionRule = append(protectionRule, rule)

		}
	}

	return protectionRule

}
