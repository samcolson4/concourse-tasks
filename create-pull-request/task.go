package main

import getPullRequests "github.com/engineerbetter/concourse-tasks/create-pull-request/lib/github"

func main() {
	ownerUsername := "engineerbetter"
	repoName := "control-tower"

	allPrs := getPullRequests.GetAllPrs(ownerUsername, repoName)
}
