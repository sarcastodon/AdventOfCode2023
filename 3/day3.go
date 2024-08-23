package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var filename = flag.String("filename", "", "file with input text")
var part = flag.String("part", "a", "AoC portion to run (expects 'a' or 'b')")

type Node struct {
	char          byte
	keepableDigit bool
}

// go run day3.go --part a --filename a.txt
func main() {
	flag.Parse()
	f, err := os.Open(*filename)
	if err != nil {
		fmt.Printf("File %s not found\n", filename)
	}

	fs := bufio.NewScanner(f)

	var nodes [][]Node

	lineLength := 0

	// Populate
	for fs.Scan() {
		line := fs.Text()
		if lineLength == 0 {
			lineLength = len(line)
		} else if lineLength != len(line) {
			fmt.Errorf("\nLINE LENGTH %d DOES NOT MATCH EXPECTED %d\n")
		}
		newRow := make([]Node, len(line))
		for i := 0; i < lineLength; i++ {
			newRow[i].char = line[i]
		}
		nodes = append(nodes, newRow)
	}
	fmt.Printf("%q", nodes)

	// Scan for parsable digits
	for j := 0; j < len(nodes); j++ {
		for i := 0; i < lineLength; i++ {
			// If not a digit, and not a '.'...
			if !strings.ContainsAny(string(nodes[j][i].char), ".0123456789") {
				fmt.Printf("\nHere's a symbol: %q\n", nodes[j][i].char)
				MarkNearbyDigitsKeepable(nodes, i, j)
			}
		}
	}

	fmt.Printf("\n\n")
	for j := 0; j < len(nodes); j++ {
		for i := 0; i < lineLength; i++ {
			if nodes[j][i].keepableDigit {
				fmt.Printf(" ")
			} else {
				fmt.Printf("%s", string(nodes[j][i].char))
			}
		}
		fmt.Printf("\n")
	}

	sum := 0
	fmt.Printf("\n")
	for j := 0; j < len(nodes); j++ {
		for i := 0; i < lineLength; i++ {
			if nodes[j][i].keepableDigit {
				var digits []byte
				for ; i < lineLength; i++ {
					if nodes[j][i].keepableDigit {
						digits = append(digits, nodes[j][i].char)
					} else {
						number, _ := strconv.Atoi(string(digits))
						fmt.Printf("| %d", number)
						sum += number
						break
					}
				}
			}
		}
	}
	fmt.Printf("\nAnswer: %d", sum)
}

func MarkNearbyDigitsKeepable(nodes [][]Node, x int, y int) {
	fmt.Printf("Calling MarkNearbyDigitsKeepable at %d, %d\n", x, y)
	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			if strings.ContainsAny(string(nodes[j][i].char), "0123456789") {
				fmt.Printf("Callng MarkContiguousDigitsKeepable at %d, %d\n", i, j)
				MarkContiguousDigitsKeepable(nodes, i, j)
			}
		}
	}
}

func MarkContiguousDigitsKeepable(nodes [][]Node, x int, y int) {
	nodes[y][x].keepableDigit = true
	for i := x; i >= 0; i-- {
		if strings.ContainsAny(string(nodes[y][i].char), "0123456789") {
			nodes[y][i].keepableDigit = true
		} else {
			break
		}
	}

	for i := x; i < len(nodes[0]); i++ {
		if strings.ContainsAny(string(nodes[y][i].char), "0123456789") {
			nodes[y][i].keepableDigit = true
		} else {
			break
		}
	}
}
