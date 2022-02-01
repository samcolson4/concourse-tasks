package main

import (
	getPullRequests "github.com/samcolson4/concourse-tasks/create-pull-request/lib/github"
	branches "github.com/samcolson4/concourse-tasks/create-pull-request/lib/branches"
	"fmt"
)

func main() {
	ownerUsername := "garethjevans"
	repoName := "create-pr"

	allPrs := getPullRequests.GetAllPrs(ownerUsername, repoName)

	// get dependency & version number that is being bumped to
	dependency := "really-useful-tool"
	newVersion := "0.3.2"

	for _, pr := range allPrs {
		pr_title := fmt.Sprintf("Bump %s to %s", dependency, newVersion)
		if pr.Title == pr_title {
			fmt.Println("Pull request already exists.")
			break
		}
	}

	branchName := fmt.Sprintf("bump-%s-to-%s", dependency, newVersion)
	branches.NewBranch(branchName)
}
