package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"unicode"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func parseForm(form string) int {
	countYes := 0
	var seen map[rune]bool = make(map[rune]bool)
	for _, r := range form {
		if unicode.IsLetter(r) && !seen[r] {
			countYes++
			seen[r] = true
		}
	}
	return countYes
}

func main() {
	dat, err := ioutil.ReadFile("adventd6-input.txt")
	check(err)
	lines := strings.Split(string(dat), "\r\n")
	countYes := 0
	var form = ""
	for _, line := range lines {
		if len(line) == 0 {
			if form != "" {
				countYes += parseForm(form)
				form = ""
			}
		} else {
			form += line + " "
		}
	}
	if form != "" {
		countYes += parseForm(form)
	}
	fmt.Println("Total yes:", countYes)
}
