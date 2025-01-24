package main

import (
	"fmt"
	"regexp"
)

func main() {
	var re = regexp.MustCompile(`(?mi)^user.*$`)
	var str = `user = username

User = username`

	for i, match := range re.FindAllString(str, -1) {
		fmt.Println(match, "found at index", i)
	}
}
