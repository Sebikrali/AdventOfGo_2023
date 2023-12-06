package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"strconv"
	"math"
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
	if !scanner.Scan() {
	    break
	}
	buffer[2] = scanner.Text()

	res = append(res, checkForAdjacentSymbols(buffer)...)
    }
    for _,v := range res {
	sum += v
    }
    fmt.Println(sum)
}

type token struct {
    index   int
    digit   int // only if a number
}

func checkForAdjacentSymbols(lines [3]string) []int {
    res := []int {}

    numbers := [3] []token{}
    symbols := [3] []token{}
    go parseLine(lines[0], &numbers[0], &symbols[0])
    go parseLine(lines[1], &numbers[1], &symbols[1])
    go parseLine(lines[2], &numbers[2], &symbols[2])

    // Idea: Iterate over the symbols of each line individually, then binary search through the number arrays
    // If a number hits then add the whole number up
    for _, v := range symbols[0] {
	
    }
    for _, v := range symbols[1] {
	
    }
    for _, v := range symbols[2] {
	
    }

    return res
}

func parseLine(line string, numbers *[]token, symbols *[]token) {
    for i, v := range line {
	if strings.Contains("0123456789", string(v)) {
	    digit, _ := strconv.Atoi(string(v))
	    *numbers = append(*numbers, token {i, digit})
	} else if strings.Contains("*#+$@-=/&", string(v)) {
	    *symbols = append(*symbols, token {i, 0})
	}
    }
}

func checkNumbers(index int, numbers *[]token) int {
    // Search first for diagonally left above
    searchIndex := index - 1
    if searchIndex < 0 {
	searchIndex++
    }

    // The symbol line is longer than the number line
    if len(*numbers) - 1 < searchIndex {
	return 0
    }

    res := 0
    low := 0
    high := len(*numbers)
    middle := (high - low) / 2 - 1
    for low <= high {
	if (*numbers)[middle].index < searchIndex {
	    low = middle
	} else if (*numbers)[middle].index > searchIndex {
	    high = middle
	} else {
	    res += getNumber((*numbers)[middle].index, middle, numbers)
	}
    }

    return 0
}

func getNumber(col, index int, numbers *[]token) int {
    proceeding := 0
    succeding := 0

    // TODO: Test that this works in a test project
    hasPredecessor := (index != 0) && ((*numbers)[index - 1].index == col - 1)
    count := 1
    for hasPredecessor {
	proceeding += int(math.Pow10(count)) * (*numbers)[index - count].digit
	hasPredecessor = (index - count - 1 != 0) && ((*numbers)[index - count].index == col - count)
    }


    return 0
}
