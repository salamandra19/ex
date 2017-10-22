package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gopl.io/ch4/github"
)

func main() {
	var less_then_month []*github.Issue
	var less_then_year []*github.Issue
	var more_then_year []*github.Issue
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d тем:\n", result.TotalCount)
	t := time.Now()
	month := t.AddDate(0, -1, 0)
	year := t.AddDate(-1, 0, 0)
	for _, item := range result.Items {
		switch {
		case item.CreatedAt.After(month):
			less_then_month = append(less_then_month, item)
		case item.CreatedAt.After(year):
			less_then_year = append(less_then_year, item)
		case item.CreatedAt.Before(year):
			more_then_year = append(more_then_year, item)
		}
	}
	fmt.Println("менее месяца:")
	for _, item := range less_then_month {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
	fmt.Println("менее года:")
	for _, item := range less_then_year {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
	fmt.Println("более года:")
	for _, item := range more_then_year {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
}
