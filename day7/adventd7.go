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

func countContainedBy(targetColor string, colors map[string]map[string]int, counts map[string]int, level int, multiplier int) int {
	totalBags := 0
	v := colors[targetColor]
	for innerKey, containedBags := range v {
		if innerKey == targetColor {
			continue
		}
		counts[innerKey] += multiplier * containedBags
		totalBags += multiplier * containedBags

		// now go down the chain
		totalBags += countContainedBy(innerKey, colors, counts, level+1, multiplier*containedBags)
		fmt.Println(strings.Repeat("\t", level), containedBags, innerKey, " =>", totalBags)

	}
	return totalBags
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

	counts2 := map[string]int{}
	totalBags := countContainedBy("shiny gold", colors, counts2, 0, 1)
	fmt.Println("Bags contained by shiny gold =", counts2, len(counts2), " Total bags =", totalBags)

}
