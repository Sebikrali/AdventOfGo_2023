package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
    // Objective: Read input, for each line take the first and last digit and concat them
    // Sum the result of all lines.
    f, err := os.Open("input")
    if err != nil {
        log.Fatal(err)
    }

    defer f.Close()

    scanner := bufio.NewScanner(f)
    sum := 0
    for scanner.Scan() {
	sum += sumFirstAndLast(scanner.Text())
    }
    fmt.Println(sum)
}

func sumFirstAndLast(input string) int {
    first := -1
    last := -1


    digits := map[string]int { "one":1,"two":2,"three":3,"four":4,"five":5,"six":6,"seven":7,"eight":8,"nine":9}

    begin := false
    str := []rune{}
    for _, v := range input {
	number, err := strconv.Atoi(string(v))
	if err != nil {
	    if begin {

	    } else if strings.Contains("otfsen", string(v)) {
		begin = true
		str = append(str, v)
	    }

	    if len(str) >= 3 {
		if res := digits[string(str)]; res != 0 {
		    last = res
		}
		checkForDigit(str, )
	    }
	} else if first == -1 {
	    first = number
	    last = number
	} else {
	    last = number
	}
    }

    return 10 * first + last
}

func checkForDigit(str []rune, state string) int {
    switch str[0] {
	case 'o':
	    	
    }

    return 0
}
