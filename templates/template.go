package main

import (
	"bufio"
	"flag"
	"os"
)

var filename = flag.String("filename", "", "file with input text")
var part = flag.String("part", "a", "AoC portion to run (expects 'a' or 'b')")

func main() {
	f, err := os.Open(*filename)

	fs := bufio.NewScanner(f)
}
