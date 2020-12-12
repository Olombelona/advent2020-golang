package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
)


// Set to true to get debug output
const verbose = true

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func extractNum(num string) int {
	v, err := strconv.Atoi(num)
	check(err)
	return v
}

func main() {
	dat, err := ioutil.ReadFile("adventd10-input.txt")
	check(err)
	lines := strings.Split(string(dat), "\r\n")
	adapters := make([]int, len(lines))

	// Use a simple insertion sort
	insertAt := 0
	for _, line := range lines {
		adapters[insertAt] = extractNum(line)
		for x := insertAt - 1; x >= 0; x-- {
			a := adapters[x]
			b := adapters[x+1]
			if a > b {
				adapters[x+1] = a
				adapters[x] = b
			} else {
				break
			}
		}
		insertAt++
	}

	// Count the jolt differences
	jolt1 := 0
	jolt2 := 0
	jolt3 := 0
	joltOthers := 0
	previous := 0
	for i := 0; i < len(adapters); i++  {
		diff := adapters[i] - previous
		previous = adapters[i]
		if diff == 1 {
			jolt1++
		} else if diff == 2 {
			jolt2++
		} else if diff == 3 {
			jolt3++
		} else {
			joltOthers++
		}
	}

	fmt.Println(adapters)
	fmt.Println("jolt1", jolt1, "jolt2", jolt2, "jolt3", jolt3, "joltOthers", joltOthers)
	// Do not forget to count the device jolt3 diff
	fmt.Println("jolt1 x jolt3+1", jolt1 * (jolt3+1))
}
