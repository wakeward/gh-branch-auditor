package rules

import (
	"encoding/json"
	"testing"

	"github.com/wakeward/gh-branch-auditor/pkg/branchprotections"
)

func Test_Allow_Force_Pushes_True(t *testing.T) {
	var data = `
{"repo_name":"example","branch":"main","is_protected":true,"protection":{"allow_force_pushes":{"enabled":true}}}
`
	var bpr *branchprotections.RepoBranchProtection

	err := json.Unmarshal([]byte(data), &bpr)
	if err != nil {
		t.Errorf("Error unmarshalling data: %v", err)
	}

	rule := AllowForcePushes(bpr)
	if rule != 1 {
		t.Errorf("Got %v expected %v", rule, 1)
	}
}

func Test_Allow_Force_Pushes_False(t *testing.T) {
	var data = `
{"repo_name":"example","branch":"main","is_protected":true,"protection":{"allow_force_pushes":{"enabled":false}}}
`
	var bpr *branchprotections.RepoBranchProtection

	err := json.Unmarshal([]byte(data), &bpr)
	if err != nil {
		t.Errorf("Error unmarshalling data: %v", err)
	}

	rule := AllowForcePushes(bpr)
	if rule != 0 {
		t.Errorf("Got %v expected %v", rule, 0)
	}
}
