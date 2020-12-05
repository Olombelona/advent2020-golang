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

func countTrees(lines []string, dx int, dy int) int {
    y := 0
    x := 0
    count := 0

    numLines := len(lines)
    
    for y < numLines {
        line := lines[y]
        if len(line) != 0 {
            mx := x % len(line)
            if (line[mx] == '#') {
                count++
            }
        }
        x += dx
        y += dy
    } 

    return count
}

/*
Right 1, down 1.
Right 3, down 1. (This is the slope you already checked.)
Right 5, down 1.
Right 7, down 1.
Right 1, down 2.
*/

func main() {
    dat, err := ioutil.ReadFile("adventd3-input.txt")
    check(err)
    lines := strings.Split(string(dat), "\r\n")
    counts := [][]int{
        {1, 1, 0},
        {3, 1, 0},
        {5, 1, 0},
        {7, 1, 0},
        {1, 2, 0}}
    mult := 1
    for _, c := range counts {
        dx := c[0]
        dy := c[1]
        count := countTrees(lines, dx, dy)
        c[2] = count
        mult *= count
        fmt.Println("Count trees dx:", dx, "dy:", dy, "count:", count)
    }
    fmt.Println(counts)
    fmt.Println("mult =", mult)
}