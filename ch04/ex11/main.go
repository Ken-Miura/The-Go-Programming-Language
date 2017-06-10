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

func main() {
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
			title := promptAndGetTitle()
			body, err := promptWithEditorAndGetBody(args[3])
			if err != nil {
				fmt.Printf("failed to get issue description: %v", err)
				return
			}
			createIssue(args[0], args[1], args[2], title, body)
			return
		} else if len(args) == 5 {
			createIssue(args[0], args[1], args[2], args[3], args[4])
			return
		} else {
			panic("This line must not be reached.\n")
		}
	} else if *operation == "edit" {
		if !(len(args) == 5 || len(args) == 6) {
			fmt.Println("usage1: " + programName + " -o edit 'owner' 'repository' 'access token' 'issue No.' 'new title' 'new description'")
			fmt.Println("ex. " + programName + ` -o edit Ken-Miura GitHub-API-Practice xxxxxxxxxx 2 ISSUE "this is new description"`)
			fmt.Println("usage2: " + programName + " -o edit 'owner' 'repository' 'access token' 'issue No.' 'editor path'")
			fmt.Println("ex. " + programName + ` -o edit Ken-Miura GitHub-API-Practice xxxxxxxxxx 2 vim`)
			return
		}
		if len(args) == 5 {
			title := promptAndGetTitle()
			body, err := promptWithEditorAndGetBody(args[4])
			if err != nil {
				fmt.Printf("failed to get issue description: %v", err)
				return
			}
			editIssue(args[0], args[1], args[2], args[3], title, body)
			return
		} else if len(args) == 6 {
			editIssue(args[0], args[1], args[2], args[3], args[4], args[5])
			return
		} else {
			panic("This line must not be reached.\n")
		}
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

func editIssue(owner, repository, accessToken, number, title, body string) {
	req, err := http.NewRequest("PATCH", gitHubAPIRoot+"/"+owner+"/"+repository+"/issues/"+number, strings.NewReader(`{"title":"`+title+`","body":"`+body+`"}`))
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

func promptAndGetTitle() string {
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
	return title
}

func promptWithEditorAndGetBody(editorPath string) (string, error) {
	fmt.Printf("Edit issue description with editor. After editing, save and close editor.\n")
	f, err := ioutil.TempFile("", "issue_util")
	if err != nil {
		return "", fmt.Errorf("failed to create temp file to edit issue description for editor use: %v", err)
	}
	tempFilePath := f.Name()
	f.Close()
	defer func() {
		if removeErr := os.Remove(tempFilePath); removeErr != nil {
			fmt.Printf("failed to delete temp file (%s): %v\n", tempFilePath, removeErr)
		}
	}()
	cmd := exec.Command(editorPath, tempFilePath)
	err = cmd.Run()
	if err != nil {
		return "", fmt.Errorf("failed to execute %s: %v", editorPath, err)
	}
	content, err := ioutil.ReadFile(tempFilePath)
	if err != nil {
		return "", fmt.Errorf("failed to read temp file (%s): %v", tempFilePath, err)
	}
	body := ""
	for _, s := range string(content) { // JSONにおいて、文字列の中で改行コードを含める際は\\nを使う。なので\r\nについて置き換える。
		if s == '\r' || s == '\n' {
			body += `\\n`
			continue
		}
		body += string(s)
	}
	return body, nil
}
