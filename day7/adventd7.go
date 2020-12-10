package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func countContainersOf(targetColor string, colors map[string]map[string]int, counts map[string]int) {
	for outerKey, v := range colors {
		for innerKey := range v {
			if innerKey == targetColor {
				counts[outerKey]++
				// leaf bag detect, now go up the chain
				countContainersOf(outerKey, colors, counts)
			}
		}
	}
}

func parseColors(color string, madeOf string, colors map[string]map[string]int) {
	if colors[color] == nil {
		colors[color] = make(map[string]int)
	}
	madeOfMap := colors[color]
	colorsToParse := strings.Split(madeOf, ",")
	var re = regexp.MustCompile(`\s*(\d+) (.+) (bag|bags)\.*`)
	for _, parseThis := range colorsToParse {
		matches := re.FindAllStringSubmatch(parseThis, -1)
		if len(matches) == 0 {
			continue
		}
		x, err := strconv.Atoi(matches[0][1])
		check(err)
		madeOfMap[matches[0][2]] += x
	}
}

func main() {
	dat, err := ioutil.ReadFile("adventd7-input.txt")
	check(err)
	lines := strings.Split(string(dat), "\r\n")
	colors := map[string]map[string]int{}
	for _, line := range lines {
		split1 := strings.Split(line, " bags contain ")
		color := split1[0]
		parseColors(color, split1[1], colors)
	}
	counts := map[string]int{}
	countContainersOf("shiny gold", colors, counts)
	fmt.Println("Bags that can contain shiny gold = ", counts, len(counts))
}
