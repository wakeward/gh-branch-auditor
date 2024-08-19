package rules

import (
	"encoding/json"
	"testing"

	"github.com/wakeward/gh-branch-auditor/pkg/branchprotections"
)

func Test_Require_Conversation_Resolution_True(t *testing.T) {
	var data = `
{"repo_name":"example","branch":"main","is_protected":true,"protection":{"required_conversation_resolution":{"enabled":true}}}
`
	var bpr *branchprotections.RepoBranchProtection

	err := json.Unmarshal([]byte(data), &bpr)
	if err != nil {
		t.Errorf("Error unmarshalling data: %v", err)
	}

	rule := RequireConversationResolution(bpr)
	if rule != 0 {
		t.Errorf("Got %v expected %v", rule, 0)
	}
}

func Test_Require_Conversation_Resolution_False(t *testing.T) {
	var data = `
{"repo_name":"example","branch":"main","is_protected":true,"protection":{"required_conversation_resolution":{"enabled":false}}}
`
	var bpr *branchprotections.RepoBranchProtection

	err := json.Unmarshal([]byte(data), &bpr)
	if err != nil {
		t.Errorf("Error unmarshalling data: %v", err)
	}

	rule := RequireConversationResolution(bpr)
	if rule != 1 {
		t.Errorf("Got %v expected %v", rule, 1)
	}
}
