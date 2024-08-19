package rules

import (
	"github.com/wakeward/gh-branch-auditor/pkg/branchprotections"
)

func BlockNewBranches(input *branchprotections.RepoBranchProtection) int {
	protectionRule := 0

	if !input.Protection.BlockCreations.GetEnabled() {
		protectionRule++
	}

	return protectionRule

}
