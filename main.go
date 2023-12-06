package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
    // Objective: Check if the numbers of a scratchcard match the winning numbers, 
    // where the first match gives 1 point and each match after that doubles the got points
    // Card 1: 41 48 83 86 17 | 83 86 6 31 17 9 48 53
    //	      winning numbers | numbers to check
    // Solution of test_input: 13

    // f, err := os.Open("input")
    f, err := os.Open("test_input")
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()

    scanner := bufio.NewScanner(f)
    sum := 0
    fmt.Println(sum)
}
