package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
)

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
	dat, err := ioutil.ReadFile("adventd8-input.txt")
	check(err)
	program := strings.Split(string(dat), "\r\n")
	accumulator := 0
	pc := 0
	numberRuns := make([]int, len(program))
	for {
		line := strings.Split(program[pc], " ")
		op := line[0]
		num := 0
		numberRuns[pc]++
		if len(line) > 1 {
			num = extractNum(line[1])
		} 
		fmt.Println(pc, op, num, "acc", accumulator)
		if (numberRuns[pc] > 1) {
			fmt.Println("Infinite loop detected at", pc, " acc =", accumulator)
			break
		}
		switch op {
		case "nop":
			pc++
		case "acc":
			accumulator += num
			pc++
		case "jmp":
			pc += num
		default:
			panic("Unknown instruction")
		}		
	}
}
