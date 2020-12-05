package main

import (
    "fmt"
    "io/ioutil"
    "strings"
    "regexp"
    "strconv"
    )

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func matchRule1(n1 int, n2 int, letter string, pwd string) bool {
    countOfLetter := strings.Count(pwd, letter)
    return countOfLetter >= n1 && countOfLetter <= n2
}

func matchRule2(n1 int, n2 int, letter string, pwd string) bool {
    var n1Match bool
    var n2Match bool
    if len(pwd) >= n1 && pwd[n1-1] == letter[0] {
        n1Match = true
    } 
    if len(pwd) >= n2 && pwd[n2-1] == letter[0] {
        n2Match = true
    } 
    return (n1Match || n2Match) && !(n1Match && n2Match)
}

func main() {
    dat, err := ioutil.ReadFile("adventd2-input.txt")
    check(err)
    lines := strings.Split(string(dat), "\r\n")
    var re = regexp.MustCompile(`(?m)(?P<min>\d+)\-(?P<max>\d+)\s+(?P<letter>.):\s+(?P<pwd>.+)`)
    totalMatching := 0
    totalMatchingNew := 0
    for _, line := range lines {
        for _, match := range re.FindAllStringSubmatch(line, -1) {
            n1,_ := strconv.Atoi(match[1])
            n2,_ := strconv.Atoi(match[2])
            letter := match[3]
            pwd := match[4]            
            if matchRule1(n1, n2, letter, pwd) {
                totalMatching++
            }
            if matchRule2(n1, n2, letter, pwd) {
                totalMatchingNew++
            }
        }
    } 
    fmt.Println("Total matching 1:", totalMatching, " out of:", len(lines))
    fmt.Println("Total matching 2:", totalMatchingNew, " out of:", len(lines))
}