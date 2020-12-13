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

var arrangements int

func countArrangements(adaptersMap map[int]bool, jolt int, arrangements map[int]int, level int) {

	if jolt == 0 {
		arrangements[0]++
		arrangements[level]++		
		if arrangements[0] % 10000000 == 0 {
			fmt.Println("Found arrangement at level", level, arrangements[0])
		}
		return
	}

	if jolt - 1 < 0 {
		fmt.Print("\vNo arrangement at level", level)
		return
	}

	for dx := 1; dx <= 3; dx++ {
		x := jolt - dx
		if adaptersMap[x] {
			countArrangements(adaptersMap, x, arrangements, level + 1)
		}
	}
}

func main() {
	dat, err := ioutil.ReadFile("adventd10-input.txt")
	check(err)
	lines := strings.Split(string(dat), "\r\n")
	adapters := make([]int, len(lines) + 2)

	// Use a simple insertion sort
	adapters[0] = 0
	insertAt := 1
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

	deviceJolt := adapters[len(adapters) - 2] + 3

	adapters[len(adapters) - 1] = deviceJolt

	adaptersMap := make(map[int]bool)

	// Count the jolt differences
	jolt1 := 0
	jolt2 := 0
	jolt3 := 0
	joltOthers := 0
	previous := 0
	adaptersMap[0] = true
	for i := 1; i < len(adapters); i++  {
		current := adapters[i]
		adaptersMap[current] = true
		diff := current - previous
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
	fmt.Println("jolt1 x jolt3", jolt1 * jolt3, "device jolt", deviceJolt)

	// All the valid combinations range from 2 to n successive adapters that add up to device jolt
	// Attempt to build trees that lead to the max device jolt
	arrangements := make(map[int]int)
	countArrangements(adaptersMap, deviceJolt - 3, arrangements, 0)

	fmt.Println("Arrangements", arrangements)
}
