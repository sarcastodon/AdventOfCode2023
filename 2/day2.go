package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var filename = flag.String("filename", "", "file with input text")
var part = flag.String("part", "a", "AoC portion to run (expects 'a' or 'b')")

func main() {
	flag.Parse()
	// Use flag to determine which file to open.
	if len(*filename) == 0 {
		log.Fatal("No filename provided.")
	}

	f, err := os.Open(*filename)
	if err != nil {
		fmt.Println("No file found.")
	}

	fs := bufio.NewScanner(f)

	if *part == "a" {
		fmt.Printf("\n\nOUTPUT: %d", a(fs))
	} else if *part == "b" {
		fmt.Printf("\n\nOUTPUT: %d", b(fs))
	}
}

var (
	gameConfig = map[string]int{
		"red":   12,
		"blue":  14,
		"green": 13,
	}
)

func a(input *bufio.Scanner) int {
	sumGames := 0
	for input.Scan() {
		possibleGame := true
		line := input.Text()
		chunks := strings.Split(line, ":")
		if len(chunks) != 2 {
			fmt.Printf("Malformed input, expected one ':' separator: %q\n", line)
		}
		game, _ := strconv.Atoi(strings.TrimPrefix(chunks[0], "Game "))
		fmt.Printf("Game number: %d\n", game)
		rounds := strings.Split(chunks[1], ";")
		for _, round := range rounds {
			fmt.Printf("Round: %s\n", round)
			moves := strings.Split(round, ",")
			for _, move := range moves {
				move = strings.TrimSpace(move)
				moveParts := strings.Split(move, " ")
				takenStones, _ := strconv.Atoi(moveParts[0])
				if takenStones > gameConfig[moveParts[1]] {
					possibleGame = false
					fmt.Printf("%s is an impossible move\n", move)
					break
				}
			}
		}
		if possibleGame == true {
			sumGames += game
			fmt.Printf("Adding game %d, new total %d", game, sumGames)
		}
	}
	return sumGames
}

func b(input *bufio.Scanner) int {
	sumGames := 0
	for input.Scan() {
		newGame := map[string]int{
			"red":   0,
			"blue":  0,
			"green": 0,
		}
		line := input.Text()
		chunks := strings.Split(line, ":")
		if len(chunks) != 2 {
			fmt.Printf("Malformed input, expected one ':' separator: %q\n", line)
		}
		game, _ := strconv.Atoi(strings.TrimPrefix(chunks[0], "Game "))
		fmt.Printf("Game number: %d\n", game)
		rounds := strings.Split(chunks[1], ";")
		for _, round := range rounds {
			fmt.Printf("Round: %s\n", round)
			moves := strings.Split(round, ",")
			for _, move := range moves {
				move = strings.TrimSpace(move)
				moveParts := strings.Split(move, " ")
				takenStones, _ := strconv.Atoi(moveParts[0])
				if takenStones > newGame[moveParts[1]] {
					newGame[moveParts[1]] = takenStones
				}
			}
		}
		sumGames += (newGame["red"] * newGame["blue"] * newGame["green"])

	}
	return sumGames
}
