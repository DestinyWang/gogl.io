// 练习 4.10 修改Issues 实例, 按照时间来输出一年内的结果
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

type User struct {
	Login               string
	id                  int
	NodeId              string          `json:"node_id"`
	AvatarURL           string          `json:"avatar_url"`
	GravatarId          string          `json:"gravatar_id"`
	URL                 string          `json:"url"`
	HTMLURL             string          `json:"html_url"`
	FollowersURL        string          `json:"followers_url"`
	FollowingURL        string          `json:"following_url"`
	GistsURL            string          `json:"gists_url"`
	StarredURL          string          `json:"starred_url"`
	SubscriptionsUrl    string          `json:"subscriptions_url"`
	OrganizationsUrl    string          `json:"organizations_url"`
	ReposUrl            string          `json:"repos_url"`
	EventsUrl           string          `json:"events_url"`
	ReceivedEventsUrl   string          `json:"received_events_url"`
	Type                string
}

type Label struct {
	id                  int
	NodeId              string          `json:"node_id"`
	URL                 string
	Name                string
	Color               string
}

type PullRequest struct {
	HTMLURL             string          `json:"html_url"`
	DiffURL             string          `json:"diff_url"`
	PATCHURL            string          `json:"patch_url"`
}

type Issue struct {
	URL                 string
	RepositoryURL       string          `json:"repository_url"`
	LabelURL            string          `json:"label_url"`
	CommentsURL         string          `json:"comments_url"`
	EventsURL           string          `json:"events_url"`
	HTMLURL             string          `json:"html_url"`
	ID                  int
	NodeId              string          `json:"node_id"`
	Number              int             `json:"number"`
	Title               string
	User                *User
	Labels              []*Label
	State               string
	//Assignee            string
	//Milestone           string
	Comments            int
	CreatedAt           time.Time       `json:"created_at"`
	UpdatedAt           time.Time       `json:"updated_at"xml:"updated_at"`
	ClosedAt            time.Time       `json:"closed_at"`
	PullRequest         *PullRequest    `json:"pull_request"`
	Body                string
	Score               float64
}

type IssuesResult struct {
	TotalCount          int             `json:"total_count"`
	IncompleteResults   bool            `json:"incomplete_results"`
	Items               []*Issue
}

const IssuesURL = "https://api.github.com/search/issues"

func SearchIssues(terms []string) (*IssuesResult, error) {
	s := strings.Join(terms, " ")   // 将参数列表用空格进行拼接 (repo:golang/go is:open json decoder)
	q := url.QueryEscape(s)             // 对 s 进行转码使之可以安全的用在 URL 查询中 (repo%3Agolang%2Fgo+is%3Aopen+json+decoder)
	resp, err := http.Get(IssuesURL + "?q=" + q)    // https://api.github.com/search/issues?q=repo%3Agolang%2Fgo+is%3Aopen+json+decoder
	if err != nil {
		return nil, err
	}
	// 必须在所有可能的分支上关闭 resp.Body
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}
	
	var result IssuesResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {  // json.NewDecoder 使用流式解码器, 可以利用它来依次从字节流中解码出多个 JSON 实体
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}

func main() {
	result, err := SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, issue := range result.Items {
		if time.Now().Year() == issue.CreatedAt.Year() {
			fmt.Printf("#%-5d %9.9s %.55s\n", issue.Number, issue.User.Login, issue.Title)
		}
	}
}
