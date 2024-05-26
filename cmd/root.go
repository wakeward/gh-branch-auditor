package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/wakeward/gh-branch-auditor/pkg/branchprotections"
)

var debug bool
var token string
var owner string
var repo string

var rootCmd = &cobra.Command{
	Use:   "gh-ba",
	Short: "A tool to audit GitHub Branch Protection Rules",
	Long:  `gh-ba is a tool to audit GitHub Branch Protection Rules.`,
	RunE: func(cmd *cobra.Command, args []string) error {

		if repo != "" {
			bpr, err := branchprotections.GetBranchProtection(owner, token, repo)
			if err != nil {
				return err
			}
			protection, _ := json.Marshal(bpr)

			fmt.Println(string(protection))
		} else {
			bpr, err := branchprotections.GetBranchProtections(owner, token)
			if err != nil {
				return err
			}
			protection, _ := json.Marshal(bpr)

			fmt.Println(string(protection))
		}

		return nil

		// TBD

		// reports, err := ruler.NewRuleset().Run(pid, mpoint)
		// if err != nil {
		// 	return err
		// }

		// var buff bytes.Buffer
		// err = output.WriteReports(format, &buff, reports)
		// if err != nil {
		// 	return err
		// }

		// out := buff.String()
		// fmt.Println(out)

		// return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "turn on debug logs")
	rootCmd.PersistentFlags().StringVarP(&token, "token", "t", "", "Set GitHub token")
	rootCmd.PersistentFlags().StringVarP(&owner, "owner", "o", "", "Set GitHub repository owner")
	rootCmd.PersistentFlags().StringVarP(&repo, "repo", "r", "", "Set GitHub repository name")
	rootCmd.MarkPersistentFlagRequired("token")
	rootCmd.MarkPersistentFlagRequired("owner")
}
