package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Starting from the beginning.")
	f, err := os.Open("test1.txt")

	if err != nil {
		fmt.Println("No file found.")
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Printf("Token: %q\n", scanner.Text())
	}

}
