package rules

import (
	"github.com/wakeward/gh-branch-auditor/pkg/branchprotections"
)

func RequireCodeOwnerReviews(input *branchprotections.RepoBranchProtection) int {
	protectionRule := 0

	if !input.Protection.RequiredPullRequestReviews.RequireCodeOwnerReviews {
		protectionRule++
	}

	return protectionRule

}
