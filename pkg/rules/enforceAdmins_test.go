package rules

import (
	"encoding/json"
	"testing"

	"github.com/wakeward/gh-branch-auditor/pkg/branchprotections"
)

func Test_Enforce_Admins_Enabled(t *testing.T) {
	var data = `
{"repo_name":"example","branch":"main","is_protected":true,"protection":{"enforce_admins":{"enabled":true}}}
`
	var bpr *branchprotections.RepoBranchProtection

	err := json.Unmarshal([]byte(data), &bpr)
	if err != nil {
		t.Errorf("Error unmarshalling data: %v", err)
	}

	rule := EnforceAdmins(bpr)
	if rule != 0 {
		t.Errorf("Got %v expected %v", rule, 0)
	}
}

func Test_Enforce_Admins_Disabled(t *testing.T) {
	var data = `
{"repo_name":"example","branch":"main","is_protected":true,"protection":{"enforce_admins":{"enabled":false}}}
`
	var bpr *branchprotections.RepoBranchProtection

	err := json.Unmarshal([]byte(data), &bpr)
	if err != nil {
		t.Errorf("Error unmarshalling data: %v", err)
	}

	rule := EnforceAdmins(bpr)
	if rule != 1 {
		t.Errorf("Got %v expected %v", rule, 1)
	}
}

func Test_Enforce_Admins_Missing(t *testing.T) {
	var data = `
{"repo_name":"example","branch":"main","is_protected":true,"protection":{}}
`
	var bpr *branchprotections.RepoBranchProtection

	err := json.Unmarshal([]byte(data), &bpr)
	if err != nil {
		t.Errorf("Error unmarshalling data: %v", err)
	}

	rule := EnforceAdmins(bpr)
	if rule != 1 {
		t.Errorf("Got %v expected %v", rule, 1)
	}
}
