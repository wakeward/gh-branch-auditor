package rules

import (
	"github.com/wakeward/gh-branch-auditor/pkg/branchprotections"
)

func AllowPushesToBranch(input *branchprotections.RepoBranchProtection) int {
	protectionRule := 0

	if !input.Protection.LockBranch.GetEnabled() {
		protectionRule++
	}

	return protectionRule

}
