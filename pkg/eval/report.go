package eval

import "github.com/wakeward/gh-branch-auditor/pkg/branchprotections"

type Reports []Report

type Report struct {
	Repo    string    `json:"repo"`
	Message string    `json:"message,omitempty"`
	Rules   []RuleRef `json:"-"`
}

type RuleRef struct {
	ID              string                                    `json:"id"`
	Reason          string                                    `json:"reason"`
	Risk            string                                    `json:"risk"`
	ProtectionRules []*branchprotections.RepoBranchProtection `json:"-"`
}
