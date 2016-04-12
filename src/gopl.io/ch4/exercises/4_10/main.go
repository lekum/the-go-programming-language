// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Issues prints a table of GitHub issues matching the search terms.
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gopl.io/ch4/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	var issuesLastMonth []string
	var issuesLastYear []string
	var issuesOlderThanAYear []string
	for _, item := range result.Items {
		s := fmt.Sprintf("#%-5d %9.9s %.55s", item.Number, item.User.Login, item.Title)
		if time.Since(item.CreatedAt) < 30*24*time.Hour {
			issuesLastMonth = append(issuesLastMonth, s)
		} else if time.Since(item.CreatedAt) < 12*30*24*time.Hour {
			issuesLastYear = append(issuesLastYear, s)
		} else {
			issuesOlderThanAYear = append(issuesOlderThanAYear, s)
		}
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	fmt.Println()

	fmt.Println("Last month:")
	fmt.Printf("%d issues\n", len(issuesLastMonth))
	for _, issue := range issuesLastMonth {
		fmt.Println(issue)
	}
	fmt.Println()

	fmt.Println("Last year:")
	fmt.Printf("%d issues\n", len(issuesLastYear))
	for _, issue := range issuesLastYear {
		fmt.Println(issue)
	}
	fmt.Println()

	fmt.Println("Previous years:")
	fmt.Printf("%d issues\n", len(issuesOlderThanAYear))
	for _, issue := range issuesOlderThanAYear {
		fmt.Println(issue)
	}
	fmt.Println()

}
