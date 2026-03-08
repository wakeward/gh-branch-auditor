package rules

import (
	"github.com/wakeward/gh-branch-auditor/pkg/branchprotections"
)

func EnforceAdmins(input *branchprotections.RepoBranchProtection) int {
	protectionRule := 0

	if input.Protection.EnforceAdmins == nil || !input.Protection.EnforceAdmins.Enabled {
		protectionRule++
	}

	return protectionRule

}
