# GitHub Branch Protection Auditor (gh-branch-auditor)

A tool for analysing GitHub Branch Protection settings on repositories

> Currently limited to the default branch where primarily branch protection rules are enforced.

- [Command line Usage](#command-line-usage)
  - [Usage Example](#usage-example)
- [To Do](#to-do)

## Command Line Usage

```bash
Usage:
  gh-ba [flags]

Flags:
      --debug          turn on debug logs
  -h, --help           help for gh-ba
  -o, --owner string   Set GitHub repository owner
  -r, --repo string    Set GitHub repository name
  -t, --token string   Set GitHub token
```

## Usage Example

Analyse all repositories for an owner

```bash
$ gh-ba -t <GITHUB_TOKEN> -o <OWNER>
```

Analyse a specific repository for an owner

```bash
$ gh-ba -t <GITHUB_TOKEN> -o <OWNER> -r <REPOSITORY>
```

## To Do

* [ ] - Define Rulesets
* [ ] - Support other branches than default
* [ ] - Create a GH-CLI extension for this