package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	words := [10]string{
		"zero",
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
	}

	fmt.Println("Starting from the beginning.")
	f, err := os.Open("day1a.text")
	//	f, err := os.Open("day1a.text")

	if err != nil {
		fmt.Println("No file found.")
	}
	fs := bufio.NewScanner(f)
	total := 0
	for fs.Scan() {
		s := fs.Text()

		earliest := len(s)
		latest := 0
		firstDigit := 0
		lastDigit := 0

		fmt.Printf("\nToken: %q\n", s)
		pattern := "0123456789"
		if first := strings.IndexAny(s, pattern); first != -1 {
			earliest = first
			firstDigit, _ = strconv.Atoi(s[first : first+1])
		}
		if last := strings.LastIndexAny(s, pattern); last != -1 {
			latest = last
			lastDigit, _ = strconv.Atoi(s[last : last+1])
		}
		// Look for string versions of numbers.

		for i := 0; i < len(words); i++ {
			if early := strings.Index(s, words[i]); early != -1 {
				if early < earliest {
					earliest = early
					firstDigit = i
				}
			}
			if late := strings.LastIndex(s, words[i]); late != -1 {
				if late > latest {
					latest = late
					lastDigit = i
				}
			}
		}
		fmt.Printf("%d (%d) and %d (%d)", earliest, firstDigit, latest, lastDigit)
		total += (firstDigit * 10) + lastDigit
		fmt.Printf("\n%d\n", total)
	}

	fmt.Printf("\n\n%d\n\n", total)
}
