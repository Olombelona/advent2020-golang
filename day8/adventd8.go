package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
)


// Set to true to get debug output
const verbose = false

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

func canTerminateProgram(program []string, pc int, accumulator int) bool {
	numberRuns := make([]int, len(program))
	lastPc := len(program) - 1

	for {
		if (pc > lastPc) {
			fmt.Println("Program terminating correctly, acc =", accumulator)
			return true
		}
		line := strings.Split(program[pc], " ")
		op := line[0]
		num := 0
		
		if len(line) > 1 {
			num = extractNum(line[1])
		} 

		numberRuns[pc]++
		if (numberRuns[pc] > 1) {
			if verbose {
				fmt.Println("  === Infinite loop detected at", pc, " acc =", accumulator)
			}
			return false
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

func solve(program []string, fix bool) {
	accumulator := 0
	pc := 0
	numberRuns := make([]int, len(program))
	lastPc := len(program) - 1

	for {
		if (pc > lastPc) {
			fmt.Println("Program terminating correctly, acc =", accumulator)
			return
		}
		line := strings.Split(program[pc], " ")
		op := line[0]
		num := 0
		
		if len(line) > 1 {
			num = extractNum(line[1])
		}
		if verbose {
			fmt.Println(pc, op, num, "acc", accumulator)
		}

		numberRuns[pc]++
		if (numberRuns[pc] > 1) {
			fmt.Println("Infinite loop detected at", pc, " acc =", accumulator)
			break
		}

		switch op {
		case "nop":
			if fix && canTerminateProgram(program, pc + num, accumulator) {
				// Change to jmp
				fmt.Println("Changing nop to jmp so code terminates at ", pc)
				return
			} 
			pc++
			
		case "acc":
			accumulator += num
			pc++
		case "jmp":
			if fix && canTerminateProgram(program, pc + 1, accumulator) {
				// Change to nop
				fmt.Println("Changing jmp to nop so code terminates at ", pc)
				return				
			}
			pc += num
			
		default:
			panic("Unknown instruction")
		}		
	}
}

func main() {
	dat, err := ioutil.ReadFile("adventd8-input.txt")
	check(err)
	program := strings.Split(string(dat), "\r\n")
	fmt.Println("== Run 1 - no fix")
	solve(program, false)
	fmt.Println("== Run 2 - fix")
	solve(program, true)
}
