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
    // Objective: Read input, for each game check if the amount of red, green or blue is possible
    // Sum the IDs of all possible Games.
    // 12 red, 13 green, 14 blue
    f, err := os.Open("input")
    if err != nil {
        log.Fatal(err)
    }

    defer f.Close()

    scanner := bufio.NewScanner(f)
    sum := 0
    for scanner.Scan() {
	sum += isPossible(scanner.Text())
	fmt.Printf("\nsum=%d\n", sum)
    }
    fmt.Println(sum)
}

func isPossible(game string) int {
    id, err := strconv.Atoi(game[5:strings.Index(game, ":")])
    fmt.Printf("id=%d, game=%s\n", id, game)
    if err != nil {
	log.Fatal(err)
    }

    tmp := game[strings.Index(game, ":")+1:]
    var index int
    for index != -1 {
	index = strings.IndexAny(tmp, "1234567890")
	if index == -1 {
	    continue;
	}
	space := index + strings.Index(tmp[index:], " ")
	number, _ := strconv.Atoi(string(tmp[index:space]))
	fmt.Printf("number=%d, string=%s,", number, tmp[space:])
	if isInvalidSet(number, tmp[space:space + 3]) {
	    return 0
	}
	tmp = tmp[index + 1:]
	fmt.Printf("tmp=%s\n", tmp)
    }
    
    return id
}

func isInvalidSet(number int, firstChar string) bool {
    fmt.Printf("func(%d, %s)\n", number, firstChar)
    switch string(firstChar[1]) {
    case "r":
	return number > 12
    case "g":
	return number > 13
    case "b":
	return number > 14
    default:
	return false
    }
}
