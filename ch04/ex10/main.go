// Copyright 2017 Ken Miura
// Issues prints a table of GitHub issues matching the search terms.
package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	result, err := SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	fmt.Println("\nissues created within 1 month (30 days)")
	for _, item := range result.Items {
		if daysAgo(item.CreatedAt) <= 30 {
			fmt.Printf("#%-5d %9.9s %.55s\n",
				item.Number, item.User.Login, item.Title)
		}
	}

	fmt.Println("\nissues created within 1 year (365 days)")
	for _, item := range result.Items {
		if daysAgo(item.CreatedAt) < 365 {
			fmt.Printf("#%-5d %9.9s %.55s\n",
				item.Number, item.User.Login, item.Title)
		}
	}

	fmt.Println("\nissues created before over 1 year (over 365 days)")
	for _, item := range result.Items {
		if daysAgo(item.CreatedAt) >= 365 {
			fmt.Printf("#%-5d %9.9s %.55s\n",
				item.Number, item.User.Login, item.Title)
		}
	}
}

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}
