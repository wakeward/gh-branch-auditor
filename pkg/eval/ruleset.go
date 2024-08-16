package eval

import (
	"sync"

	"github.com/wakeward/gh-branch-auditor/pkg/branchprotections"
	"github.com/wakeward/gh-branch-auditor/pkg/rules"
)

type Ruleset struct {
	Rules []Rule
}

func NewRuleset() *Ruleset {
	list := make([]Rule, 0)

	isProtected := Rule{
		Predicate: rules.IsProtected,
		ID:        "IsProtected",
		Reason:    "Branch Protection is Disabled",
		Risk:      "High",
	}
	list = append(list, isProtected)

	allowForcePushes := Rule{
		Predicate: rules.AllowForcePushes,
		ID:        "AllowForcePushes",
		Reason:    "Force push overwrites current branch with another",
		Risk:      "High",
	}
	list = append(list, allowForcePushes)

	return &Ruleset{
		Rules: list,
	}
}

func (rs *Ruleset) Run(bpr []*branchprotections.RepoBranchProtection) ([]Report, error) {
	reports := make([]Report, 0)
	report := Report{
		Repo:  bpr[0].RepoName,
		Rules: make([]RuleRef, 0),
	}

	// run rules in parallel
	ch := make(chan RuleRef, len(rs.Rules))
	var wg sync.WaitGroup
	for _, rule := range rs.Rules {
		wg.Add(1)
		go eval(bpr, rule, ch, &wg)
	}
	wg.Wait()
	close(ch)

	// collect results
	// var appliedRules int
	for ruleRef := range ch {
		if ruleRef.Risk != "" {
			report.Rules = append(report.Rules, ruleRef)
		}
	}

	reports = append(reports, report)
	return reports, nil
}

func eval(bpr []*branchprotections.RepoBranchProtection, rule Rule, ch chan RuleRef, wg *sync.WaitGroup) {
	defer wg.Done()

	protectionRules := rule.Eval(bpr)

	result := RuleRef{
		ID:              rule.ID,
		Risk:            rule.Risk,
		Reason:          rule.Reason,
		ProtectionRules: protectionRules,
	}

	ch <- result
}
