package branchprotections

import (
	"context"
	"fmt"

	"github.com/google/go-github/v62/github"
)

func AuthClient(token string) (client *github.Client) {
	return github.NewClient(nil).WithAuthToken(token)
}

// get all branch protections

func GetBranchProtections(owner string, token string) (bpr *github.Protection, err error) {

	client := AuthClient(token)

	// get all repositories for an owner

	repos, _, err := client.Repositories.ListByUser(context.Background(), owner, nil)

	for _, repo := range repos {

		// get all branch protections
		bpr, _, err = client.Repositories.GetBranchProtection(context.Background(), owner, *repo.Name, *repo.DefaultBranch)

		if err != nil {
			return nil, fmt.Errorf("error getting branch protections: %w", err)
		}

	}

	return bpr, err
}
