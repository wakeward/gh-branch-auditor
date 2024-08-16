package rules

import (
	"github.com/wakeward/gh-branch-auditor/pkg/branchprotections"
)

func AllowForcePushes(input []*branchprotections.RepoBranchProtection) []*branchprotections.RepoBranchProtection {
	protectionRule := make([]*branchprotections.RepoBranchProtection, 0)

	for _, rule := range input {
		if rule.Protection.AllowForcePushes.Enabled {
			protectionRule = append(protectionRule, rule)
		}
	}

	return protectionRule

}
