// Copyright 2017 Ken Miura
package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"
)

const gitHubAPIRoot = "https://api.github.com/repos"

var operation = flag.String("o", "list", "issue operation for the repository: list, create, edit or close")

// TODO エディタ連携 エディタを起動した後にそのエディタに入力された文字列はどうやって取得すればよい？？？
func main() {
	//createIssue("Ken-Miura", "GitHub-API-Practice", "997a0cc0c949edd2a288736294d2375ea55917c0", "test6", "")
	flag.Parse()
	if !(*operation == "list" || *operation == "create" || *operation == "edit" || *operation == "close") {
		fmt.Println("Operation must be list, create, edit or close.")
		return
	}

	programName := os.Args[0]
	args := flag.Args()
	if *operation == "list" {
		if len(args) != 2 {
			fmt.Println("usage: " + programName + " [-o list] 'owner' 'repository'")
			fmt.Println("ex1. " + programName + " Ken-Miura GitHub-API-Practice")
			fmt.Println("ex2. " + programName + " -o list Ken-Miura GitHub-API-Practice")
			return
		}
		listIssues(args[0], args[1])
		return
	} else if *operation == "create" {
		if !(len(args) == 4 || len(args) == 5) {
			fmt.Println("usage1: " + programName + " -o create 'owner' 'repository' 'access token' 'title' 'body'")
			fmt.Println("ex. " + programName + ` -o create Ken-Miura GitHub-API-Practice xxxxxxxxxx issue1 "this is description""`)
			fmt.Println("usage2: " + programName + " -o create 'owner' 'repository' 'access token' 'editor path'")
			fmt.Println("ex. " + programName + " -o create Ken-Miura GitHub-API-Practice xxxxxxxxxx vim")
			fmt.Println("caution: Save contents as UTF-8 If you use editor.")
			return
		}
		if len(args) == 4 {
			// title取得処理
			fmt.Printf("issue title: ")
			scanner := bufio.NewScanner(os.Stdin)
			title := ""
			for scanner.Scan() {
				title = scanner.Text()
				if title != "" {
					break
				}
				fmt.Println("Enter issue title.")
				fmt.Printf("issue title: ")
			}
			// ここまでtitle取得処理

			// body取得処理
			fmt.Printf("Edit issue description with editor. After editing, save and close editor.\n")
			f, err := ioutil.TempFile("", "issue_util")
			if err != nil {
				fmt.Printf("failed to create temp file to edit issue description for editor use: %v", err)
				return
			}
			tempFilePath := f.Name()
			f.Close()
			defer func() {
				err = os.Remove(tempFilePath)
				if err != nil {
					fmt.Printf("failed to delete temp file (%s): %v", tempFilePath, err)
				}
			}()
			cmd := exec.Command(args[3], tempFilePath)
			err = cmd.Run()
			if err != nil {
				fmt.Printf("failed to execute %s: %v", args[2], err)
				return
			}
			content, err := ioutil.ReadFile(tempFilePath)
			if err != nil {
				fmt.Printf("failed to read temp file (%s): %v", tempFilePath, err)
				return
			}
			// ここまでbody取得処理

			createIssue(args[0], args[1], args[2], title, string(content))
			return
		} else if len(args) == 5 {
			createIssue(args[0], args[1], args[2], args[3], args[4])
			return
		} else {
			panic("This line must not be reached.\n")
		}
	} else if *operation == "edit" {
		if len(args) != 5 {
			fmt.Println("usage: " + programName + " -o edit 'owner' 'repository' 'access token' 'issue No.' 'new title'")
			fmt.Println("ex. " + programName + " -o edit Ken-Miura GitHub-API-Practice xxxxxxxxxx 2 ISSUE")
			return
		}
		editIssue(args[0], args[1], args[2], args[3], args[4])
		return
	} else if *operation == "close" {
		if len(args) != 4 {
			fmt.Println("usage: " + programName + " -o close 'owner' 'repository' 'access token' 'issue No.'")
			fmt.Println("ex. " + programName + " -o close Ken-Miura GitHub-API-Practice xxxxxxxxxx 1")
			return
		}
		closeIssue(args[0], args[1], args[2], args[3])
		return
	} else {
		panic("This line must not be reached.\n")
	}
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

func listIssues(owner, repository string) {
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
		fmt.Println(err)
		return
	}

	for _, item := range issues {
		fmt.Printf("No.%d state: %s title: %.55s\n", item.Number, item.State, item.Title)
	}
}

func createIssue(owner, repository, accessToken, title, body string) {
	req, err := http.NewRequest("POST", gitHubAPIRoot+"/"+owner+"/"+repository+"/issues", strings.NewReader(`{"title":"`+title+`","body":"`+body+`"}`))
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Set("Authorization", "token "+accessToken)
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 201 {
		var message struct{ Message string }
		if err := json.NewDecoder(resp.Body).Decode(&message); err != nil {
			fmt.Printf("failed to decode response body as json: %v", err)
		}
		fmt.Printf("failed to request. status code: %v. reason: %v", resp.StatusCode, message)
		return
	}
}

func editIssue(owner, repository, accessToken, number, title string) {
	req, err := http.NewRequest("PATCH", gitHubAPIRoot+"/"+owner+"/"+repository+"/issues/"+number, strings.NewReader(`{"title":"`+title+`"}`))
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Set("Authorization", "token "+accessToken)
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var message struct{ Message string }
		if err := json.NewDecoder(resp.Body).Decode(&message); err != nil {
			fmt.Printf("failed to decode response body as json: %v", err)
		}
		fmt.Printf("failed to request. status code: %v. reason: %v", resp.StatusCode, message)
		return
	}
}

func closeIssue(owner, repository, accessToken, number string) {
	req, err := http.NewRequest("PATCH", gitHubAPIRoot+"/"+owner+"/"+repository+"/issues/"+number, strings.NewReader(`{"state":"closed"}`))
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Set("Authorization", "token "+accessToken)
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var message struct{ Message string }
		if err := json.NewDecoder(resp.Body).Decode(&message); err != nil {
			fmt.Printf("failed to decode response body as json: %v", err)
		}
		fmt.Printf("failed to request. status code: %v. reason: %v", resp.StatusCode, message)
		return
	}
}
