package main

import (
	"fmt"
	"time"
)

func main() {

	// Build Start
	start := time.Date(2023, 6, 8, 0, 0, 0, 0, time.UTC)
	fmt.Println("Start Date:", start.Format("2006-01-02"))

	// Run Ignore Weekends
	duration := 14 // 2 weeks duration
	end := addDaysIgnoreWeekends(start, duration)

	fmt.Println("End Date (skipping weekends):", end.Format("2006-01-02"))
}

func addDaysIgnoreWeekends(start time.Time, days int) time.Time {
	for i := 0; i < days; {
		start = start.Add(24 * time.Hour)
		// if the day is Saturday or Sunday, continue to the next day without counting it
		if start.Weekday() != time.Saturday && start.Weekday() != time.Sunday {
			i++
		}
	}
	return start
}
