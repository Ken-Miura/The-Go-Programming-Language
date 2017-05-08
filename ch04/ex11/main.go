// Copyright 2017 Ken Mirua
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"
)

const gitHubAPIRoot = "https://api.github.com/repos"

var operation = flag.String("o", "list", "issue operation for the repository: list, create or edit")

func main() {
	flag.Parse()
	args := flag.Args()
	if !(*operation == "list" || *operation == "create" || *operation == "edit") {
		fmt.Println("Operation must be list, create or edit.")
		return
	}

	if *operation == "list" {
		if len(args) != 2 {
			fmt.Println("usage: " + os.Args[0] + " [-o list] 'owner' 'repository'")
			fmt.Println("ex1. " + os.Args[0] + " -o list Ken-Miura GitHub-API-Practice")
			return
		}
		resp, err := http.Get(gitHubAPIRoot + "/" + args[0] + "/" + args[1] + "/issues")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			fmt.Println(resp.StatusCode)
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
