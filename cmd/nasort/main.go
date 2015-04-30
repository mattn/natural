package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/mattn/natural"
	"os"
)

var f = flag.Bool("f", "fold", true)

func main() {
	flag.Parse()

	var l lines
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		l = append(l, scanner.Text())
	}
	if *f {
		natural.SortCase(l)
	} else {
		natural.Sort(l)
	}
	for _, line := range l {
		fmt.Println(line)
	}
}
