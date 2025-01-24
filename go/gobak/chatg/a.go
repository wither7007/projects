package main

import (
	"fmt"
	"time"
)

func getTimeThreshold(weeks int) time.Time {
	// Subtract n weeks from the current time
	return time.Now().AddDate(0, 0, -weeks*7)
}

func main() {
	weeks := 7
	timeThreshold := getTimeThreshold(weeks)
	fmt.Println(timeThreshold)
	fmt.Println(time.Now())
}
