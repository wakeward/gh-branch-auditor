package rules

import (
	"github.com/wakeward/gh-branch-auditor/pkg/branchprotections"
)

func RequireSignedCommits(input *branchprotections.RepoBranchProtection) int {
	protectionRule := 0

	if !input.Protection.RequiredSignatures.GetEnabled() {
		protectionRule++
	}

	return protectionRule

}
