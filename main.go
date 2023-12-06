package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
    // Objective: To complicated, might add later

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
