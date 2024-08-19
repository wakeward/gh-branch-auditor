package rules

import (
	"encoding/json"
	"testing"

	"github.com/wakeward/gh-branch-auditor/pkg/branchprotections"
)

func Test_Allow_Pushes_True(t *testing.T) {
	var data = `
{"repo_name":"example","branch":"main","is_protected":true,"protection":{"lock_branch":{"enabled":true}}}
`
	var bpr *branchprotections.RepoBranchProtection

	err := json.Unmarshal([]byte(data), &bpr)
	if err != nil {
		t.Errorf("Error unmarshalling data: %v", err)
	}

	rule := AllowPushesToBranch(bpr)
	if rule != 0 {
		t.Errorf("Got %v expected %v", rule, 0)
	}
}

func Test_Allow_Pushes_False(t *testing.T) {
	var data = `
{"repo_name":"example","branch":"main","is_protected":true,"protection":{"lock_branch":{"enabled":false}}}
`
	var bpr *branchprotections.RepoBranchProtection

	err := json.Unmarshal([]byte(data), &bpr)
	if err != nil {
		t.Errorf("Error unmarshalling data: %v", err)
	}

	rule := AllowPushesToBranch(bpr)
	if rule != 1 {
		t.Errorf("Got %v expected %v", rule, 1)
	}
}
