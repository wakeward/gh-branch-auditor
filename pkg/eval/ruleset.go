package eval

import (
	"sync"
	"time"

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

func (rs *Ruleset) Run(bpr []*branchprotections.RepoBranchProtection) (Reports, error) {
	// reports := make([]Report, 0)

	reports := Reports{
		Title:  "GitHub Branch Auditor",
		Verson: "0.1",
		Date:   time.Now().Format("2006-01-02T15:04:05"),
		Report: make([]Report, 0),
	}

	for _, repo := range bpr {

		report := Report{
			Repo:   repo.RepoName,
			Branch: repo.Branch,
			Rules:  make([]RuleRef, 0),
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
		if repo.IsProtected {
			for ruleRef := range ch {
				for _, protection := range ruleRef.ProtectionRules {
					if protection != nil {
						report.Rules = append(report.Rules, ruleRef)
					}
				}
			}
			reports.Report = append(reports.Report, report)

		}

	}
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
