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
	fmt.Printf("%d тем:\n", result.TotalCount)

	var less_than_month []*github.Issue
	var less_than_year []*github.Issue
	var more_than_year []*github.Issue

	t := time.Now()
	month := t.AddDate(0, -1, 0)
	year := t.AddDate(-1, 0, 0)
	for _, item := range result.Items {
		switch {
		case item.CreatedAt.After(month):
			less_than_month = append(less_than_month, item)
		case item.CreatedAt.After(year):
			less_than_year = append(less_than_year, item)
		default:
			more_than_year = append(more_than_year, item)
		}
	}

	printIssue("менее месяца:", less_than_month)
	printIssue("менее года:", less_than_year)
	printIssue("более года:", more_than_year)
}

func printIssue(s string, period []*github.Issue) {
	fmt.Println(s)
	for _, item := range period {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
}
