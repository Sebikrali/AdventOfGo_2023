package main

import (
	"bufio"
	"fmt"
	"log"
	"strconv"
	"os"
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
    for _, v := range input {
	number, err := strconv.Atoi(string(v))
	if err != nil {
	    continue
	} else if first == -1 {
	    first = number
	    last = number
	} else {
	    last = number
	}
    }

    return 10 * first + last
}
