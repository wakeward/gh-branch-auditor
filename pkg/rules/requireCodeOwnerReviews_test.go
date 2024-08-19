package rules

import (
	"encoding/json"
	"testing"

	"github.com/wakeward/gh-branch-auditor/pkg/branchprotections"
)

func Test_Require_CodeOwner_Reviews_True(t *testing.T) {
	var data = `
{"repo_name":"example","branch":"main","is_protected":true,"protection":{"required_pull_request_reviews":{"require_code_owner_reviews":true}}}
`
	var bpr *branchprotections.RepoBranchProtection

	err := json.Unmarshal([]byte(data), &bpr)
	if err != nil {
		t.Errorf("Error unmarshalling data: %v", err)
	}

	rule := RequireCodeOwnerReviews(bpr)
	if rule != 0 {
		t.Errorf("Got %v expected %v", rule, 0)
	}
}

func Test_Require_CodeOwner_Reviews_False(t *testing.T) {
	var data = `
{"repo_name":"example","branch":"main","is_protected":true,"protection":{"required_pull_request_reviews":{"require_code_owner_reviews":false}}}
`
	var bpr *branchprotections.RepoBranchProtection

	err := json.Unmarshal([]byte(data), &bpr)
	if err != nil {
		t.Errorf("Error unmarshalling data: %v", err)
	}

	rule := RequireCodeOwnerReviews(bpr)
	if rule != 1 {
		t.Errorf("Got %v expected %v", rule, 1)
	}
}
