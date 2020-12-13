package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"sync"
)

// Set to true to get debug output
const verbose = true
const reportStep = 1000

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

func countArrangements(adaptersMap map[int]bool, jolt int, arrangements map[int]int, level int, id int) {

	if jolt == 0 {
		arrangements[0]++
		arrangements[level]++
		a := arrangements[0]
		if (a % reportStep) == 0 {
			fmt.Println(id, "\tFound arrangement at level", level, a)
		}
		return
	}

	if jolt-1 < 0 {
		fmt.Println("No arrangement at level", level)
		return
	}

	for dx := 1; dx <= 3; dx++ {
		x := jolt - dx
		if adaptersMap[x] {
			countArrangements(adaptersMap, x, arrangements, level+1, id)
		}
	}
}

func worker(adaptersMap map[int]bool, jolt int, arrangements map[int]int, level int, lastMap int, wg *sync.WaitGroup) map[int]int {
	defer wg.Done()
	id := lastMap*10000 + jolt
	fmt.Printf("=>>> Starting worker %d\n", id)
	countArrangements(adaptersMap, jolt, arrangements, level+1, id)
	fmt.Printf("<<<= Worker %d done\n", id)
	return arrangements
}

func mergeArrangements(arrangements map[int]map[int]int) map[int]int {
	m := make(map[int]int)
	for _, ar := range arrangements {
		for k, v := range ar {
			m[k] += v
		}
	}
	return m
}

func main() {
	dat, err := ioutil.ReadFile("adventd10-input-short.txt")
	check(err)
	lines := strings.Split(string(dat), "\n")
	adapters := make([]int, len(lines)+2)

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

	deviceJolt := adapters[len(adapters)-2] + 3

	adapters[len(adapters)-1] = deviceJolt

	adaptersMap := make(map[int]bool)

	// Count the jolt differences
	jolt1 := 0
	jolt2 := 0
	jolt3 := 0
	joltOthers := 0
	previous := 0
	adaptersMap[0] = true
	for i := 1; i < len(adapters); i++ {
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
	fmt.Println("jolt1 x jolt3", jolt1*jolt3, "device jolt", deviceJolt)

	// All the valid combinations range from 2 to n successive adapters that add up to device jolt
	// Attempt to build trees that lead to the max device jolt
	arrangements := make(map[int]map[int]int)
	lastMap := 0

	var wg sync.WaitGroup

	for j1 := 1; j1 <= 3; j1++ {
		xj1 := deviceJolt - 3 - j1
		if adaptersMap[xj1] {
			for j11 := 1; j11 <= 3; j11++ {
				xj11 := xj1 - j11
				if adaptersMap[xj11] {
					wg.Add(1)
					ar := make(map[int]int)
					arrangements[lastMap] = ar
					lastMap++
					go worker(adaptersMap, xj11, ar, 2, lastMap, &wg)
				}
			}
		}
	}

	wg.Wait()

	fmt.Println("Arrangements", mergeArrangements(arrangements))
}
