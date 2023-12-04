package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
    // Objective: Read input, sum all numbers that are adjacent to a symbol like *,#, basically everything that isn't a number
    // and '.' doesn't count
    // 467..114..
    // ...*......
    // ..35..633.
    // Solution of test_input: 4361
    // f, err := os.Open("input")
    f, err := os.Open("test_input")
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()

    scanner := bufio.NewScanner(f)
    sum := 0

    buffer := [3]string{}
    for i := 0; i < 3; i++ {
	scanner.Scan()
	buffer[i] = scanner.Text()
    }
    res := checkForAdjacentSymbols(buffer)

    for scanner.Scan() {
	buffer[0] = buffer[2]
	buffer[1] = scanner.Text()
	if scanner.Scan() {
	    buffer[2] = scanner.Text()
	}
	res = append(res, checkForAdjacentSymbols(buffer)...)
    }
    for _,v := range res {
	sum += v
    }
    fmt.Println(sum)
}

type token struct {
    index	int
    numOrSym	int // 0 for number, 1 for symbol
    digit	int // only if a number
}

func checkForAdjacentSymbols(lines [3]string) []int {
    // Idea: Go through the lines and add the indexes of numbers and symbols to a slice
    //	     one slice for each line

    parsedLines := [3] []token{}
    // TODO: Parse each line in a separat go routine

    return nil
}

func parseLine(line string, result *[]token) {

}
