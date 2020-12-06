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

func parseForm2(form string) int {
	countAllYes := 0

	var lines = strings.Split(strings.Trim(form, "\n"), "\n")
	var counts map[rune]int = make(map[rune]int)
	for _, line := range lines {
		var seen map[rune]bool = make(map[rune]bool)
		for _, r := range line {
			if unicode.IsLetter(r) && !seen[r] {
				if _, hasCount := counts[r]; hasCount {
					counts[r]++
				} else {
					counts[r] = 1
				}
				seen[r] = true
			}
		}
	}
	numPeople := len(lines)
	for _, count := range counts {
		if count == numPeople {
			countAllYes++
		}
	}
	return countAllYes
}

func main() {
	dat, err := ioutil.ReadFile("adventd6-input.txt")
	check(err)
	lines := strings.Split(string(dat), "\r\n")
	countYes := 0
	countAllYes := 0
	var form = ""
	for _, line := range lines {
		if len(line) == 0 {
			if form != "" {
				countYes += parseForm(form)
				countAllYes += parseForm2(form)
				form = ""
			}
		} else {
			form += line + "\n"
		}
	}
	if form != "" {
		countYes += parseForm(form)
		countAllYes += parseForm2(form)
	}
	fmt.Println("Total yes:", countYes)
	fmt.Println("Total all yes:", countAllYes)
}
