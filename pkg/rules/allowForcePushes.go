package rules

import (
	"github.com/wakeward/gh-branch-auditor/pkg/branchprotections"
)

func AllowForcePushes(input *branchprotections.RepoBranchProtection) int {
	protectionRule := 0

	if input.Protection.AllowForcePushes.Enabled {
		protectionRule++
	}

	return protectionRule

}
