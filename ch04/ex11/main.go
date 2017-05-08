// Copyright 2017 Ken Mirua
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

const gitHubAPIRoot = "https://api.github.com/repos"

var operation = flag.String("o", "list", "issue operation for the repository: list, create, edit or close")

func main() {
	flag.Parse()
	args := flag.Args()
	if !(*operation == "list" || *operation == "create" || *operation == "edit" || *operation == "close") {
		fmt.Println("Operation must be list, create, edit or close.")
		return
	}

	if *operation == "list" {
		if len(args) != 2 {
			fmt.Println("usage: " + os.Args[0] + " [-o list] 'owner' 'repository'")
			fmt.Println("ex1. " + os.Args[0] + " Ken-Miura GitHub-API-Practice")
			fmt.Println("ex2. " + os.Args[0] + " -o list Ken-Miura GitHub-API-Practice")
			return
		}
		resp, err := http.Get(gitHubAPIRoot + "/" + args[0] + "/" + args[1] + "/issues")
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
	} else if *operation == "create" {
		if len(args) != 4 {
			fmt.Println("usage: " + os.Args[0] + " -o create 'owner' 'repository' 'access token' 'title'")
			fmt.Println("ex. " + os.Args[0] + " -o create Ken-Miura GitHub-API-Practice xxxxxxxxxx issue1")
			return
		}

		req, err := http.NewRequest("POST", gitHubAPIRoot+"/"+args[0]+"/"+args[1]+"/issues", strings.NewReader(`{"title":"`+args[3]+`"}`))
		if err != nil {
			fmt.Println(err)
			return
		}
		req.Header.Set("Authorization", "token "+args[2])
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
				fmt.Printf("failed to decode response body as json. error: %v", err)
			}
			fmt.Printf("failed to request. status code: %v. reason: %v", resp.StatusCode, message)
			return
		}
	} else if *operation == "edit" {
		if len(args) != 5 {
			fmt.Println("usage: " + os.Args[0] + " -o edit 'owner' 'repository' 'access token' 'issue No.' 'new title'")
			fmt.Println("ex. " + os.Args[0] + " -o edit Ken-Miura GitHub-API-Practice xxxxxxxxxx 2 ISSUE")
			return
		}

		req, err := http.NewRequest("PATCH", gitHubAPIRoot+"/"+args[0]+"/"+args[1]+"/issues/"+args[3], strings.NewReader(`{"title":"`+args[4]+`"}`))
		if err != nil {
			fmt.Println(err)
			return
		}
		req.Header.Set("Authorization", "token "+args[2])
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
				fmt.Printf("failed to decode response body as json. error: %v", err)
			}
			fmt.Printf("failed to request. status code: %v. reason: %v", resp.StatusCode, message)
			return
		}
	} else if *operation == "close" {
		if len(args) != 4 {
			fmt.Println("usage: " + os.Args[0] + " -o close 'owner' 'repository' 'access token' 'issue No.'")
			fmt.Println("ex. " + os.Args[0] + " -o close Ken-Miura GitHub-API-Practice xxxxxxxxxx 1")
			return
		}
		req, err := http.NewRequest("PATCH", gitHubAPIRoot+"/"+args[0]+"/"+args[1]+"/issues/"+args[3], strings.NewReader(`{"state":"closed"}`))
		if err != nil {
			fmt.Println(err)
			return
		}
		req.Header.Set("Authorization", "token "+args[2])
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
				fmt.Printf("failed to decode response body as json. error: %v", err)
			}
			fmt.Printf("failed to request. status code: %v. reason: %v", resp.StatusCode, message)
			return
		}
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
