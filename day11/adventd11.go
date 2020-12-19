package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)


// Set to true to get debug output
const verbose = false
const smallInput = false
const floor byte = 0
const empty byte = 1
const occupied byte = 2

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func adjacentOccupiedCount(seats [][]byte, row int, seat int) int {
	coords := [][2]int {
		{-1, -1},{0, -1}, {1, -1},
		{-1, 0}, /* {0, 0} */ {1, 0},
		{-1, 1}, {0,1}, {1, 1}}
	countOccupied := 0
	maxX := len(seats[row])-1
	maxY := len(seats)-1
	for _, coord := range coords {
		x := seat + coord[0]
		y := row + coord[1]
		if verbose && (row == 8) && (seat == 9) {
			var v byte = 99
			if x >= 0 && x <= maxX  && y >= 0 && y <= maxY {
				v = seats[y][x]
			}			
			fmt.Println(y, x, v)
		}
		if x >= 0 && x <= maxX  && y >= 0 && y <= maxY && seats[y][x] == occupied {
			countOccupied++
		}
	}
	return countOccupied
}

func seatsToString(seats [][]byte) string {
	var sb strings.Builder
	var runes = []rune{'.', 'L', '#'}
	for y, row := range seats {
		sb.WriteString(fmt.Sprintf("%d\t", y))
		for _, seat := range row {
			sb.WriteRune(runes[seat])
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

func simulateSeating(seats [][]byte) ([][]byte, int, int) {
	newSeats := make([][]byte, len(seats))
	changes := 0
	countOccupied := 0
	for y, row := range seats {
		newSeats[y] = make([]byte, len(row))
		for x, seat := range row {
			if seat == floor {
				newSeats[y][x] = floor
			} else if seat == empty && adjacentOccupiedCount(seats, y, x) == 0 {
				newSeats[y][x] = occupied
				changes++
				countOccupied++
			} else if seat == occupied {
				if adjacentOccupiedCount(seats, y, x) >= 4 {
					newSeats[y][x] = empty
					changes++
				} else {
					newSeats[y][x] = occupied
					countOccupied++
				}
			} else {
				newSeats[y][x] = seat
			}
		}
	}
	return newSeats, changes, countOccupied
}

func main() {
	input := "adventd11-input.txt"
	if smallInput {
		input = "adventd11-input-small.txt"
	}
	dat, err := ioutil.ReadFile(input)
	check(err)
	lines := strings.Split(string(dat), "\r\n")
	seats := make([][]byte, len(lines));
	for i, line := range lines {		
		seats[i] = make([]byte, len(line));
		for j, seat := range line {
			if (seat == 'L') {
				seats[i][j] = empty
			} else if (seat == '#') {
				seats[i][j] = occupied
			} else {
				seats[i][j] = floor
			}
		}
	}

	continueSimulation := true
	var newSeats [][]byte = seats
	countOccupied := 0
	changes := 0
	for continueSimulation {
		newSeats, changes, countOccupied = simulateSeating(newSeats)
		continueSimulation = changes > 0
		if verbose {
			fmt.Print(seatsToString(newSeats))
		}
		fmt.Println("changes", changes, "occupied", countOccupied)
	}	
	
}
