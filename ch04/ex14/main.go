// Copyright 2017 Ken Mirua
package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
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
	handler := func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}

		var owner, repository string
		for k, v := range r.Form {
			if k == "owner" {
				owner = v[0]
			} else if k == "repository" {
				repository = v[0]
			}
		}
		reportBug(w, owner, repository)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func reportBug(out io.Writer, owner, repository string) {
	resp, err := http.Get(gitHubAPIRoot + "/" + owner + "/" + repository + "/issues")
	if err != nil {
		fmt.Fprintln(out, err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Fprintf(out, "failed to request. status code: %v\n", resp.StatusCode)
		return
	}

	var issues []Issue
	if err := json.NewDecoder(resp.Body).Decode(&issues); err != nil {
		fmt.Fprintf(out, "failed to decode response body as json. error: %v\n", err)
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
		fmt.Fprintf(out, "failed to write bug report to specified output stream. error: %v\n", err)
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
