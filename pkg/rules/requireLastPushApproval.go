package rules

import (
	"github.com/wakeward/gh-branch-auditor/pkg/branchprotections"
)

func RequireLastPushApproval(input *branchprotections.RepoBranchProtection) int {
	protectionRule := 0

	if !input.Protection.RequiredPullRequestReviews.RequireLastPushApproval {
		protectionRule++
	}

	return protectionRule

}
