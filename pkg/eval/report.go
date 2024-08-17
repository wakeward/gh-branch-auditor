package eval

import "github.com/wakeward/gh-branch-auditor/pkg/branchprotections"

type Reports struct {
	Title  string   `json:"title"`
	Verson string   `json:"version"`
	Date   string   `json:"date"`
	Report []Report `json:"report"`
}

type Report struct {
	Repo   string    `json:"repo"`
	Branch string    `json:"branch"`
	Rules  []RuleRef `json:"rules"`
}

type RuleRef struct {
	ID              string                                    `json:"id"`
	Reason          string                                    `json:"reason"`
	Risk            string                                    `json:"risk"`
	ProtectionRules []*branchprotections.RepoBranchProtection `json:"-"`
}
