package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/mattn/natural"
	"os"
)

var f = flag.Bool("f", true, "fold case")

func main() {
	flag.Parse()

	var lines []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if *f {
		natural.SortCase(lines)
	} else {
		natural.Sort(lines)
	}
	for _, line := range lines {
		fmt.Println(line)
	}
}
