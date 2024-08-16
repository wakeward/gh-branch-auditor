package branchprotections

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/google/go-github/v62/github"
)

type RepoBranchProtection struct {
	RepoName    string            `json:"repo_name"`
	Branch      string            `json:"branch"`
	IsProtected bool              `json:"is_protected"`
	Protection  github.Protection `json:"protection,omitempty"`
}

func AuthClient(token string) (client *github.Client) {
	return github.NewClient(nil).WithAuthToken(token)
}

func GetBranchProtections(owner string, token string, repoName string) (rbp []*RepoBranchProtection, err error) {

	var repos []*github.Repository
	var allRepos []*RepoBranchProtection

	client := AuthClient(token)

	if repoName != "" {
		repo, _, err := client.Repositories.Get(context.Background(), owner, repoName)

		if err != nil {
			return nil, fmt.Errorf("error getting repository: %w", err)
		}

		repos = append(repos, repo)

	} else {
		repos, _, err = client.Repositories.ListByUser(context.Background(), owner, nil)

		if err != nil {
			return nil, fmt.Errorf("error getting repositories: %w", err)
		}
	}

	for _, repo := range repos {
		// get default branch
		branch, _, err := client.Repositories.GetBranch(context.Background(), owner, *repo.Name, *repo.DefaultBranch, 1)
		if err != nil {
			return nil, fmt.Errorf("error getting default branch: %w", err)
		}

		if *branch.Protected {
			bpr, _, err := client.Repositories.GetBranchProtection(context.Background(), owner, *repo.Name, *repo.DefaultBranch)
			if err != nil {
				return nil, fmt.Errorf("error getting branch protection: %w", err)
			}

			repoList := &RepoBranchProtection{
				RepoName:    *repo.Name,
				Branch:      *repo.DefaultBranch,
				IsProtected: *branch.Protected,
				Protection:  *bpr,
			}

			raw_output, _ := json.Marshal(repoList)
			fmt.Println(string(raw_output))

			allRepos = append(allRepos, repoList)

		} else {
			repoList := &RepoBranchProtection{
				RepoName:    *repo.Name,
				Branch:      *repo.DefaultBranch,
				IsProtected: *branch.Protected,
			}

			allRepos = append(allRepos, repoList)
		}

	}

	return allRepos, err
}
