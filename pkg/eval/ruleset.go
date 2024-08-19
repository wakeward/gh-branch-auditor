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

func ProtectedRuleset() *Ruleset {
	list := make([]Rule, 0)

	allowDeletions := Rule{
		Predicate: rules.AllowDeletions,
		ID:        "GH-BP-003",
		Rule:      "Allow deletions",
		Risk:      "Protected branch configuration can removed",
		Severity:  "High",
	}

	list = append(list, allowDeletions)

	allowForcePushes := Rule{
		Predicate: rules.AllowForcePushes,
		ID:        "GH-BP-002",
		Rule:      "Allow force pushes",
		Risk:      "Force push overwrites current branch with another",
		Severity:  "High",
	}
	list = append(list, allowForcePushes)

	allowPushesToBranch := Rule{
		Predicate: rules.AllowPushesToBranch,
		ID:        "GH-BP-011",
		Rule:      "Lock branch",
		Risk:      "Push directly to branch is allowed for collaborators and teams",
		Severity:  "High",
	}
	list = append(list, allowPushesToBranch)

	blockNewBranches := Rule{
		Predicate: rules.BlockNewBranches,
		ID:        "GH-BP-006",
		Rule:      "Restrict who can push to matching branches",
		Risk:      "New branches can be created by any user",
		Severity:  "Low",
	}
	list = append(list, blockNewBranches)

	dismissStaleReviews := Rule{
		Predicate: rules.DismissStaleReviews,
		ID:        "GH-BP-005",
		Rule:      "Dismiss stale pull request approvals when new commits are pushed",
		Risk:      "New commits does not require a code review",
		Severity:  "Medium",
	}
	list = append(list, dismissStaleReviews)

	requireCodeOwnerReview := Rule{
		Predicate: rules.RequireCodeOwnerReviews,
		ID:        "GH-BP-004",
		Rule:      "Require review from Code Owners",
		Risk:      "Code owner pull request review is not required",
		Severity:  "High",
	}
	list = append(list, requireCodeOwnerReview)

	requireConversationResolution := Rule{
		Predicate: rules.RequireConversationResolution,
		ID:        "GH-BP-007",
		Rule:      "Require conversation resolution before merging",
		Risk:      "Not all comments need to be resolved before pull request is merged.",
		Severity:  "Low",
	}

	list = append(list, requireConversationResolution)

	requireSignedCommits := Rule{
		Predicate: rules.RequireSignedCommits,
		ID:        "GH-BP-012",
		Rule:      "Require signed commits",
		Risk:      "Not all commits are not signed",
		Severity:  "Low",
	}
	list = append(list, requireSignedCommits)

	requireLinearHistory := Rule{
		Predicate: rules.RequireLinearHistory,
		ID:        "GH-BP-008",
		Rule:      "Require linear history",
		Risk:      "Merge commits can be pushed to the branch",
		Severity:  "Medium",
	}
	list = append(list, requireLinearHistory)

	requireLastPushApproval := Rule{
		Predicate: rules.RequireLastPushApproval,
		ID:        "GH-BP-009",
		Rule:      "Require approval of the most recent reviewable push",
		Risk:      "Last user to push changes can approve the pull request",
		Severity:  "Medium",
	}

	list = append(list, requireLastPushApproval)

	strictStatusChecks := Rule{
		Predicate: rules.StrictStatusChecks,
		ID:        "GH-BP-010",
		Rule:      "Require status checks to pass before merging",
		Risk:      "Branches do not need to up to date before merging.",
		Severity:  "Medium",
	}

	list = append(list, strictStatusChecks)

	return &Ruleset{
		Rules: list,
	}
}

func UnprotectedRuleset() *Ruleset {
	list := make([]Rule, 0)

	isProtected := Rule{
		Predicate: rules.IsProtected,
		ID:        "GH-BP-001",
		Rule:      "Branch protection applied to listed branch",
		Risk:      "Branch Protection is Disabled",
		Severity:  "High",
	}
	list = append(list, isProtected)

	return &Ruleset{
		Rules: list,
	}
}

func (rs *Ruleset) Run(bpr []*branchprotections.RepoBranchProtection) (Reports, error) {
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
			go eval(repo, rule, ch, &wg)
		}
		wg.Wait()
		close(ch)

		// collect results
		for ruleRef := range ch {
			if ruleRef.ProtectionRules > 0 {
				report.Rules = append(report.Rules, ruleRef)
			}
		}
		reports.Report = append(reports.Report, report)
	}

	return reports, nil
}

func eval(bpr *branchprotections.RepoBranchProtection, rule Rule, ch chan RuleRef, wg *sync.WaitGroup) {
	defer wg.Done()

	protectionRules := rule.Eval(bpr)

	result := RuleRef{
		ID:              rule.ID,
		Rule:            rule.Rule,
		Risk:            rule.Risk,
		Severity:        rule.Severity,
		ProtectionRules: protectionRules,
	}

	ch <- result
}
