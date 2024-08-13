package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Starting from the beginning.")
	f, err := os.Open("day1a.text")

	if err != nil {
		fmt.Println("No file found.")
	}
	fs := bufio.NewScanner(f)
	total := 0
	for fs.Scan() {
		s := fs.Text()
		fmt.Printf("\nToken: %q\n", s)
		pattern := "0123456789"
		first := strings.IndexAny(s, pattern)
		firstNum, _ := strconv.Atoi(s[first : first+1])
		last := strings.LastIndexAny(s, pattern)
		lastNum, _ := strconv.Atoi(s[last : last+1])
		fmt.Printf("%d (%d) and %d (%d)", first, firstNum, last, lastNum)
		total += (firstNum * 10) + lastNum
		fmt.Printf("\n\n%d\n\n", total)
	}
	fmt.Printf("\n\n%d\n\n", total)
}
