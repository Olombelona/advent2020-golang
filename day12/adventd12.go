package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
	"math"
)

// Set to true to get debug output
const verbose = false

const east = 0
const north = 1
const west = 2
const south = 3

const slotLeft = 0
const slotRight = 1


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

func extractNumf(num string) float64 {
	v, err := strconv.ParseFloat(num, 64)
	check(err)
	return v
}


func turnShip(facing int, turnBy int, slot int) int {
	//      N=1
	//  W=2     E=0
	//      S=3
	directions := [][]int{ {north, south}, {west, east}, {south, north}, {east, west} }
	newFacing := facing	
	steps := (turnBy / 90)
	
	for s := 1; s <= steps; s++ {
		newFacing = directions[newFacing][slot]
	}

	return newFacing
}

func rotateWaypoint(wx float64, wy float64, degrees float64) (float64, float64) {
	// source: https://stackoverflow.com/a/34374437
	radians := degrees * (math.Pi/180.0)

	// Assume origin (0, 0)
	qx := math.Cos(radians) * wx - math.Sin(radians) * wy
	qy := math.Sin(radians) * wx + math.Cos(radians) * wy

	return  math.Round(qx), math.Round(qy)
}

func solve2(input string) {
	fmt.Println("=============== Solving part 2", input)
	dat, err := ioutil.ReadFile(input)
	check(err)
	lines := strings.Split(string(dat), "\r\n")
	
	mx := 0.0
	my := 0.0
	wx := 10.0
	wy := 1.0
	for _, line := range lines {		
		dir := line[0]
		v := extractNumf(line[1:])
		switch dir {
			case 'N':
				wy += v
			case 'S':
				wy -= v
			case 'E':
				wx += v
			case 'W':
				wx -= v
			case 'L':
				wx, wy = rotateWaypoint(wx, wy, v)
			case 'R':
				// Angle sign is different than part 1
				wx, wy = rotateWaypoint(wx, wy, -v)
			case 'F':
				mx += v * wx
				my += v * wy
			default:
				panic("unknow instruction")
		}
	}
	fmt.Println("mx", mx, "my", my, "wx", wx, "wy", wy, "distance",  math.Abs(mx) + math.Abs(my))
}

func solve(input string) {
	fmt.Println("=============== Solving part 1", input)
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
				facing = turnShip(facing, v, slotLeft)
			case 'R':
				facing = turnShip(facing, v, slotRight)
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
	}
	fmt.Println("facing", facing, "ew", ewPosition, "ns", nsPosition, "distance", abs(ewPosition) + abs(nsPosition))
}

func main() {
	solve("adventd12-input.txt")
	solve("adventd12-input-small.txt")
	solve2("adventd12-input.txt")
	solve2("adventd12-input-small.txt")
}
