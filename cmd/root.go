package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/wakeward/gh-branch-auditor/pkg/branchprotections"
)

var debug bool
var token string
var owner string
var outputLocation string

var rootCmd = &cobra.Command{
	Use:   "gh-branch-auditor",
	Short: "A tool to audit GitHub branch protection rules",
	Long:  `gh-branch-auditor is a tool to audit GitHub branch protection rules.`,
	RunE: func(cmd *cobra.Command, args []string) error {

		bps, err := branchprotections.GetBranchProtections(owner, token)
		if err != nil {
			return err
		}

		fmt.Println(bps)

		return nil

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
	rootCmd.PersistentFlags().StringVarP(&outputLocation, "output", "l", "", "Set output location")
}
