package rules

import (
	"github.com/wakeward/gh-branch-auditor/pkg/branchprotections"
)

func IsProtected(input *branchprotections.RepoBranchProtection) int {
	protectionRule := 0

	if !input.IsProtected {
		protectionRule++
	}

	return protectionRule

}
