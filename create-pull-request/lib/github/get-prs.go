package getPullRequests

import (
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"time"
)

type PullRequests []struct {
	URL      string `json:"url"`
	ID       int    `json:"id"`
	Number   int    `json:"number"`
	State    string `json:"state"`
	Title    string `json:"title"`
	User     struct {
		Login             string `json:"login"`
		HTMLURL           string `json:"html_url"`
	} `json:"user"`
	CreatedAt          time.Time     `json:"created_at"`
	UpdatedAt          time.Time     `json:"updated_at"`
	ClosedAt           interface{}   `json:"closed_at"`
	MergedAt           interface{}   `json:"merged_at"`
	MergeCommitSha     string        `json:"merge_commit_sha"`
	Assignee           interface{}   `json:"assignee"`
	Assignees          []interface{} `json:"assignees"`
	RequestedReviewers []interface{} `json:"requested_reviewers"`
	RequestedTeams     []interface{} `json:"requested_teams"`
	Labels             []struct {
		Name        string `json:"name"`
	} `json:"labels"`
	Head              struct {
		Label string `json:"label"`
		Ref   string `json:"ref"`
		Sha   string `json:"sha"`
		User  struct {
			Login             string `json:"login"`
			HTMLURL           string `json:"html_url"`
		} `json:"user"`
		Repo struct {
			ID       int    `json:"id"`
			Name     string `json:"name"`
			FullName string `json:"full_name"`
			HTMLURL          string      `json:"html_url"`
			DefaultBranch string   `json:"default_branch"`
		} `json:"repo"`
	} `json:"head"`
	Base struct {
		Label string `json:"label"`
		Ref   string `json:"ref"`
		Sha   string `json:"sha"`
		User  struct {
			Login             string `json:"login"`
			HTMLURL           string `json:"html_url"`
		} `json:"user"`
		Repo struct {
			ID       int    `json:"id"`
			Name     string `json:"name"`
			FullName string `json:"full_name"`
			Owner    struct {
				Login             string `json:"login"`
				HTMLURL           string `json:"html_url"`
			} `json:"owner"`
			HTMLURL          string      `json:"html_url"`
		} `json:"repo"`
	} `json:"base"`
}

func GetAllPrs(ownerUsername string, repoName string) PullRequests {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/pulls", ownerUsername, repoName)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	bodyString := string(body)

	if bodyString == "[]" {
		fmt.Printf("No pull requests on %s/%s.", ownerUsername, repoName)
		return nil
	}

	var prs PullRequests
	err = json.Unmarshal(body, &prs)
	if err != nil {
		log.Fatalln(err)
	}

	return prs
}
