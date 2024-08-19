package rules

import (
	"github.com/wakeward/gh-branch-auditor/pkg/branchprotections"
)

func RequireConversationResolution(input *branchprotections.RepoBranchProtection) int {
	protectionRule := 0

	if !input.Protection.RequiredConversationResolution.Enabled {
		protectionRule++
	}

	return protectionRule

}
