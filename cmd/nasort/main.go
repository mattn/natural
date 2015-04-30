package main

import (
	"bufio"
	"fmt"
	"github.com/mattn/natural"
	"os"
	"sort"
)

type lines []string

func (l lines) Len() int {
	return len(l)
}

func (l lines) Less(i, j int) bool {
	return natural.NaturalComp(l[i], l[j]) < 0
}

func (l lines) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func main() {
	var l lines
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		l = append(l, scanner.Text())
	}
	sort.Sort(l)
	for _, line := range l {
		fmt.Println(line)
	}
}
