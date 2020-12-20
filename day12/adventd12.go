package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
)

// Set to true to get debug output
const verbose = false

const east = 0
const north = 1
const west = 2
const south = 3

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func abs(x int) int {
	if x >= 0 { return x }
	return -x
}

func extractNum(num string) int {
	v, err := strconv.Atoi(num)
	check(err)
	return v
}

func turnShip(facing int, turnBy int) int {
	//      N=1
	//  W=2     E=0
	//      S=3
	directions := [][]int{ {north, south}, {west, east}, {south, north}, {east, west} }
	newFacing := facing	
	steps := (turnBy / 90)
	if steps > 0 {
		for s := 1; s <= steps; s++ {
			newFacing = directions[newFacing][1]
		}
	} else {
		for s := -1; s >= steps; s-- {
			newFacing = directions[newFacing][0]
		}
	}

	return newFacing
}

func solve(input string) {
	fmt.Println("=============== Solving", input)
	dat, err := ioutil.ReadFile(input)
	check(err)
	lines := strings.Split(string(dat), "\r\n")
	
	facing := east
	ewPosition := 0
	nsPosition := 0
	for _, line := range lines {		
		dir := line[0]
		v := extractNum(line[1:])
		switch dir {
			case 'N':
				nsPosition += v
			case 'S':
				nsPosition -= v
			case 'E':
				ewPosition += v
			case 'W':
				ewPosition -= v
			case 'L':
				facing = turnShip(facing, -v)
			case 'R':
				facing = turnShip(facing, v)
			case 'F':
				if facing == east {
					ewPosition += v
				} else if facing == north {
					nsPosition += v
				} else if facing == west {
					ewPosition -= v
				} else if facing == south {
					nsPosition -= v
				}
			default:
				panic("unknow instruction")
		}
		fmt.Println(line, "facing", facing, "ew", ewPosition, "ns", nsPosition, "distance", abs(ewPosition) + abs(nsPosition))
	}
}

func main() {
	solve("adventd12-input.txt")
	solve("adventd12-input-small.txt")
}
