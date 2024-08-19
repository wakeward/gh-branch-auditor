# GitHub Branch Protection Auditor (gh-branch-auditor)

A tool for analysing GitHub Branch Protection settings on repositories

> Currently limited to the default branch where primarily branch protection rules are enforced.

- [Required Permissions](#required-permissions)
- [Command line Usage](#command-line-usage)
  - [Usage Example](#usage-example)
- [Rulesets](#rulesets)

## Required Permissions

To assess the branch protection settings for a repository, the user person access token **must have** ["Administration" repository permissions (read)](https://docs.github.com/en/rest/branches/branch-protection?apiVersion=2022-11-28#get-branch-protection)

## Command Line Usage

```bash
gh-branch-auditor is a tool to audit GitHub Branch Protection Rules.

Usage:
  gh-branch-auditor [flags]

Flags:
      --debug           turn on debug logs
  -f, --format string   Set output format (cli, json) (default "json")
  -h, --help            help for gh-branch-auditor
  -o, --owner string    Set GitHub repository owner
  -r, --repo string     Set GitHub repository name
  -t, --token string    Set GitHub token
```

## Usage Example

Analyse all repositories for an owner

```bash
$ gh-branch-auditor -o <OWNER>
```

Analyse a specific repository for an owner

```bash
$ gh-branch-auditor -o <OWNER> -r <REPOSITORY>
```

By default, the GitHub Branch Auditor will look for a `gh` CLI token, a token can be specified with `-t` flag.

```bash
$ gh-branch-auditor -t <GH_TOKEN> -o <OWNER> -r <REPOSITORY>
```

## Rulesets

The following rulesets are analysed by the gh-branch-auditor.

| RuleSet ID | Branch Rule | Risk | Severity |
|-----|-----|-----|-----|
| GH-BP-001 | Branch protection applied to listed branch | Branch Protections is Disabled. | High |
| GH-BP-002 | Allow force pushes | Force push overwrites current branch with another. | High |
| GH-BP-003 | Allow deletions | Protected branch configuration can removed. | High |
| GH-BP-004 | Require review from Code Owners | Code owner pull request review is not required. | High |
| GH-BP-005 | Dismiss stale pull request approvals when new commits are pushed | New commits does not require a code review. | Medium |
| GH-BP-006 | Restrict who can push to matching branches | New branches can be created by any user. | Low |
| GH-BP-007 | Require conversation resolution before merging | Not all comments need to be resolved before pull request is merged. | Low |
| GH-BP-008 | Require linear history | Merge commits can be pushed to the branch. | Medium |
| GH-BP-009 | Require approval of the most recent reviewable push | Last user to push changes can approve the pull request. | Medium |
| GH-BP-010 | Require status checks to pass before merging | Branches do not need to up to date before merging. | Medium |
| GH-BP-011 | Lock branch | Push directly to branch is allowed for collaborators and teams. | High |
| GH-BP-012 | Require signed commits | Not all commits are not signed | Low |

> Note: There are no risks that are considered critical as user must be authenticated as collaborator or team to push to the repository.

### Limitations

The current limitations for the tool are as follows:

* Analysis is limited to GitHub SaaS (github.com)
* The assessment is only against the default branch (where primarily branch protection rules are enforced)
* Currently only json output is supported