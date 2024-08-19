package rules

import (
	"github.com/wakeward/gh-branch-auditor/pkg/branchprotections"
)

func RequireLinearHistory(input *branchprotections.RepoBranchProtection) int {
	protectionRule := 0

	if !input.Protection.RequireLinearHistory.Enabled {
		protectionRule++
	}

	return protectionRule

}
