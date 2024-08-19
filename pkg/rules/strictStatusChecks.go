package rules

import (
	"github.com/wakeward/gh-branch-auditor/pkg/branchprotections"
)

func StrictStatusChecks(input *branchprotections.RepoBranchProtection) int {
	protectionRule := 0

	if !input.Protection.RequiredStatusChecks.Strict {
		protectionRule++
	}

	return protectionRule

}
