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

	var inMonthIssues, inYearIssues, otherIssues []*github.Issue
	for _, item := range result.Items {
		if item.CreatedAt.Month() == time.Now().Month() {
			inMonthIssues = append(inMonthIssues, item)
		} else if item.CreatedAt.Year() == time.Now().Year() {
			inYearIssues = append(inYearIssues, item)
		} else {
			otherIssues = append(otherIssues, item)
		}
	}

	fmt.Printf("%d issues less than a month old:\n", len(inMonthIssues))
	for _, issue := range inMonthIssues {
		fmt.Printf("#%-5d %9.9s %.55s \n",
			issue.Number, issue.User.Login, issue.Title)
	}

	fmt.Printf("\n%d issues less than a year old:\n", len(inYearIssues))
	for _, issue := range inYearIssues {
		fmt.Printf("#%-5d %9.9s %.55s \n",
			issue.Number, issue.User.Login, issue.Title)
	}

	fmt.Printf("\n%d issues more than a year old:\n", len(otherIssues))
	for _, issue := range otherIssues {
		fmt.Printf("#%-5d %9.9s %.55s \n",
			issue.Number, issue.User.Login, issue.Title)
	}

	fmt.Printf("\nTotal: %d issues\n", result.TotalCount)
}
