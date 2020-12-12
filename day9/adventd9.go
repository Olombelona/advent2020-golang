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

func isValidNumber(numbers []int, at int, preambleSize int) bool {
	sortedPredecessors := make([]int, preambleSize)
	insertAt := 0
	maxCheckAt := preambleSize - 1
	n := numbers[at]	
	comps := 0
	// Insertion sort during the first run
	for i := at - preambleSize; i < at; i++ {
		sortedPredecessors[insertAt] = numbers[i]
		for x := insertAt - 1; x >= 0; x-- {
			a := sortedPredecessors[x]
			b := sortedPredecessors[x+1]
			comps++
			if a + b == n {
				// Worth checking early
				if verbose {
					fmt.Println("Comps at ", at, " - early validation = ", comps, "out of", 25*25)
				}
				return true
			}
			if a > b {
				// Insert until sorted
				sortedPredecessors[x+1] = a
				sortedPredecessors[x] = b
			} else {
				// Correctly sorted, stop
				break
			}
		}
		if sortedPredecessors[insertAt] > n {
			// No need to check anything above as the sum is going to be automatically greater
			maxCheckAt = insertAt
		}
		insertAt++
	}
	// Second pass will use the sorted array to validate
	for i := 0; i <= maxCheckAt; i++ {
		for j := i + 1; j <= maxCheckAt; j++ {
			comps++
			if sortedPredecessors[i] + sortedPredecessors[j] == n {
				if verbose {
					fmt.Println("Comps at ", at, " = ", comps, "out of", 25*25)
				}
				return true
			}
		}
	}
	if verbose {
		fmt.Println("Comps at ", at, " = ", comps, "out of", 25*25)
	}
	return false
}

func main() {
	dat, err := ioutil.ReadFile("adventd9-input.txt")
	check(err)
	lines := strings.Split(string(dat), "\r\n")
	numbers := make([]int, len(lines))

	for i, line := range lines {
		numbers[i] = extractNum(line)
	}

	for i := 25; i < len(numbers); i++ {
		if !isValidNumber(numbers, i, 25) {
			fmt.Println("First invalid number at pos", i, " val =", numbers[i])
			return
		}
	}

	fmt.Println(len(numbers))
}
