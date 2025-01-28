package cmd

import (
	"bytes"
	"fmt"
	"os"

	"github.com/cli/go-gh/v2/pkg/auth"
	"github.com/spf13/cobra"
	"github.com/wakeward/gh-branch-auditor/pkg/branchprotections"
	"github.com/wakeward/gh-branch-auditor/pkg/eval"
	"github.com/wakeward/gh-branch-auditor/pkg/output"
)

var debug bool
var token string
var owner string
var repo string
var format string

var rootCmd = &cobra.Command{
	Use:   "gh-branch-auditor",
	Short: "A tool to audit GitHub Branch Protection Rules",
	Long:  `gh-branch-auditor is a tool to audit GitHub Branch Protection Rules.`,
	RunE: func(cmd *cobra.Command, args []string) error {

		if token == "" {
			t, _ := auth.TokenForHost("github.com")
			token = t
		}

		bpr, err := branchprotections.GetBranchProtections(owner, token, repo)
		if err != nil {
			return err
		}

		reports, err := IsProtected(bpr)
		if err != nil {
			return err
		}

		var buff bytes.Buffer
		err = output.WriteReports(format, &buff, reports)
		if err != nil {
			return err
		}

		out := buff.String()
		fmt.Println(out)

		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "turn on debug logs")
	rootCmd.PersistentFlags().StringVarP(&token, "token", "t", "", "Set GitHub token")
	rootCmd.PersistentFlags().StringVarP(&owner, "owner", "o", "", "Set GitHub repository owner")
	rootCmd.PersistentFlags().StringVarP(&repo, "repo", "r", "", "Set GitHub repository name")
	rootCmd.PersistentFlags().StringVarP(&format, "format", "f", "json", "Set output format (cli, json)")
	rootCmd.MarkPersistentFlagRequired("owner")
}

func IsProtected(bpr []*branchprotections.RepoBranchProtection) (reports eval.Reports, err error) {
	for _, repo := range bpr {
		if repo.IsProtected {
			reports, err = eval.ProtectedRuleset().Run(bpr)
			if err != nil {
				return reports, err
			}
		} else if !repo.IsProtected {
			reports, err = eval.UnprotectedRuleset().Run(bpr)
			if err != nil {
				return reports, err
			}
		}
	}
	return reports, nil
}
