package main

import (
    "fmt"
    "io/ioutil"
    "strings"
    )

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func binaryFind(line string, startLow int, startHigh int, stepStart int, stepEnd int) int {
    low := startLow
    high := startHigh    
    // fmt.Println("start", low, high)
    for i := stepStart; i <= stepEnd; i++ {
        dir := line[i]        
        switch dir {
        case 'F', 'L': 
            // take the lower half
            high = low + (((high - low) + 1) / 2) - 1
            if (i == stepEnd) { return low }
        case 'B', 'R':
            // take the upper half
            low = low + (((high - low) + 1) / 2)           
            if (i == stepEnd) { return high }
        }
        // fmt.Println(string(dir), low, high)
    }
    return low
}

func main() {
    dat, err := ioutil.ReadFile("adventd5-input.txt")
    check(err)
    highestSeatNumber := 0
    lines := strings.Split(string(dat), "\r\n")
    for _, line := range lines {
        row :=binaryFind(line, 0, 127, 0, 6)
        col :=binaryFind(line, 0, 7, 7, 9)
        seat := row * 8 + col
        if seat > highestSeatNumber { highestSeatNumber = seat }
        fmt.Println(line, row, col, seat)
    } 
    fmt.Println("Highest seat number", highestSeatNumber)
}