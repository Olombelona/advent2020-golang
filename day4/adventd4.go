package main

import (
    "fmt"
    "io/ioutil"
	"strings"
	"strconv"
	"regexp"
    )

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func validDate(str string, min int, max int) bool {
	year, err := strconv.Atoi(str)
	if err == nil && year >= min && year <= max { return true }
	return false
}

func validHeight(hgt string) bool {
// hgt (Height) - a number followed by either cm or in:
// If cm, the number must be at least 150 and at most 193.
// If in, the number must be at least 59 and at most 76.
	var re = regexp.MustCompile(`(?m)(\d+)(cm|in)`)
	matches := re.FindAllStringSubmatch(hgt, -1)

	if len(matches) != 1 { return false }

	if len(matches[0]) != 3 { return false }

	x, err := strconv.Atoi(matches[0][1])

	if (err != nil) {return false}

	if matches[0][2] == "cm" {
		return x >= 150 && x <= 193
	}

	if matches[0][2] == "in" {
		return x >= 59 && x <= 76
	}

	return true
}

// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
func validHairColor(str string) bool {
	var re = regexp.MustCompile(`(?m)\#[a-f0-9]{6}`)
	return re.MatchString(str)
}

// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
func validEyeColor(str string) bool {
	var re = regexp.MustCompile(`(?m)(amb|blu|brn|gry|grn|hzl|oth)`)
	return re.MatchString(str)
}

// pid (Passport ID) - a nine-digit number, including leading zeroes.
func validPassportID(str string) bool {
	var re = regexp.MustCompile(`(?m)[0-9]{9}`)
	return re.MatchString(str)
}

func isValidPassport(passport string) bool {
	if (passport == "") {return false}

	var m map[string]string = make(map[string]string);
	for _, field := range strings.Split(passport, " ") {
		pair := strings.Split(field, ":")
		if (len(pair) == 2) {
			m[pair[0]] = pair[1]
		}
	}
	/*
	byr (Birth Year)
	iyr (Issue Year)
	eyr (Expiration Year)
	hgt (Height)
	hcl (Hair Color)
	ecl (Eye Color)
	pid (Passport ID)
	cid (Country ID)
	*/
	byrValue, byr := m["byr"]
	iyrValue, iyr := m["iyr"]
	eyrValue, eyr := m["eyr"]
	hgtValue, hgt := m["hgt"]
	hclValue, hcl := m["hcl"]
	eclValue, ecl := m["ecl"]
	pidValue, pid := m["pid"]

	if !(byr && iyr && eyr && hgt && hcl && ecl && pid) { return false }

	if !validDate(byrValue, 1920, 2002) { return false }
	if !validDate(iyrValue, 2010, 2020) { return false }
	if !validDate(eyrValue, 2020, 2030) { return false }
	if !validHeight(hgtValue) { return false }
	if !validHairColor(hclValue) { return false }
	if !validEyeColor(eclValue) { return false }
	if !validPassportID(pidValue) { return false }

// byr (Birth Year) - four digits; at least 1920 and at most 2002.
// iyr (Issue Year) - four digits; at least 2010 and at most 2020.
// eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
// hgt (Height) - a number followed by either cm or in:
// If cm, the number must be at least 150 and at most 193.
// If in, the number must be at least 59 and at most 76.
// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
// pid (Passport ID) - a nine-digit number, including leading zeroes.
// cid (Country ID) - ignored, missing or not.	
	return true
}

func main() {
    dat, err := ioutil.ReadFile("adventd4-input.txt")
    check(err)
	lines := strings.Split(string(dat), "\r\n")
	countValid := 0
	var passport = ""
	for _,line := range lines {
		if len(line) == 0 {
			if isValidPassport(passport) {
				countValid++
				fmt.Println(">>", passport)
			}
			passport = ""
		} else {
			passport += line + " "
		}
	}
	if passport!= "" && isValidPassport(passport) {
		countValid++
		fmt.Println("??", passport)
	}
	fmt.Println("Valid passports:", countValid)
}