package rules

import (
	"github.com/wakeward/gh-branch-auditor/pkg/branchprotections"
)

func AllowDeletions(input *branchprotections.RepoBranchProtection) int {
	protectionRule := 0

	if input.Protection.AllowDeletions.Enabled {
		protectionRule++
	}

	return protectionRule

}
