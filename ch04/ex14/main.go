// Copyright 2017 Ken Mirua
package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"time"
)

const gitHubAPIRoot = "https://api.github.com/repos"

var bugReport = template.Must(template.New("bugReport").Funcs(template.FuncMap{"hasMilestone": hasMilestone}).Parse(`
<h1>{{.TotalCount}} bugs</h1>
<table>
<tr style='text-align: left'>
  <th>#</th>
  <th>State</th>
  <th>User</th>
  <th>Title</th>
  <th>Milestone</th>
</tr>
{{range .Items}}
<tr>
  <td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
  <td>{{.State}}</td>
  <td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
  <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
  <td><a href='{{(.Milestone|hasMilestone).HTMLURL}}'>{{(.Milestone|hasMilestone).Title}}</a></td>
</tr>
{{end}}
</table>
`))

func hasMilestone(milestone *Milestone) *Milestone {
	if milestone == nil {
		return &Milestone{"no associated milestone", ""}
	} else {
		return milestone
	}
}

func main() {
	reportBug(os.Stdout, "Ken-Miura", "GitHub-API-Practice")
}

func reportBug(out io.Writer, owner, repository string) {
	resp, err := http.Get(gitHubAPIRoot + "/" + owner + "/" + repository + "/issues")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("failed to request. status code: %v\n", resp.StatusCode)
		return
	}

	var issues []Issue
	if err := json.NewDecoder(resp.Body).Decode(&issues); err != nil {
		fmt.Printf("failed to decode response body as json. error: %v\n", err)
		return
	}

	bugCount := 0
	bugIssues := make([]Issue, 0)
NextIssue:
	for _, issue := range issues {
		for _, label := range *issue.Labels {
			if label.Name == "bug" {
				bugIssues = append(bugIssues, issue)
				bugCount++
				continue NextIssue
			}
		}
	}
	bugs := BugList{bugCount, bugIssues}
	if err := bugReport.Execute(out, bugs); err != nil {
		fmt.Printf("failed to write bug report to response body. error: %v\n", err)
		return
	}
}

type BugList struct {
	TotalCount int
	Items      []Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	Milestone *Milestone
	User      *User
	Labels    *[]Label
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

type Label struct {
	Name string
}

type Milestone struct {
	Title   string
	HTMLURL string `json:"html_url"`
}
