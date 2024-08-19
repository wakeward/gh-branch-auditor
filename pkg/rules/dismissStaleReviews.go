package rules

import (
	"github.com/wakeward/gh-branch-auditor/pkg/branchprotections"
)

func DismissStaleReviews(input *branchprotections.RepoBranchProtection) int {
	protectionRule := 0

	if !input.Protection.RequiredPullRequestReviews.DismissStaleReviews {
		protectionRule++
	}

	return protectionRule

}
