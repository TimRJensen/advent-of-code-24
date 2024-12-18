package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func task1And2(lists [][]int, fault int) int {
	sum := 0
	for _, lst := range lists {
		isum := 0
		for i, n := 0, 0; i < len(lst)-1; i++ {
			delta := lst[i+1] - lst[i]
			isum += delta
			if delta > 0 && delta < 4 {
				continue
			}
			if n == fault {
				isum = len(lst) * 3
				break
			}
			n++
		}

		jsum := 0
		for j, n := len(lst)-1, 0; j > 0; j-- {
			delta := lst[j-1] - lst[j]
			jsum += delta
			if delta > 0 && delta < 4 {
				continue
			}
			if n == fault {
				jsum = len(lst) * 3
				break
			}
			n++
		}

		if (isum >= len(lst)-1 && isum <= (len(lst)-1)*3) || (jsum >= len(lst)-1 && jsum <= (len(lst)-1)*3) {
			sum++
		}
	}

	return sum
}

func parse(path string) [][]int {
	buff, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	ns := make([][]int, 0, 32)
	for _, line := range strings.Split(string(buff), "\n") {
		ns = append(ns, make([]int, 0, 32))
		for _, n := range strings.Fields(strings.TrimSpace(line)) {
			if n, ok := strconv.Atoi(n); ok == nil {
				ns[len(ns)-1] = append(ns[len(ns)-1], n)
			}
		}
	}

	return ns
}

func main() {
	lists := parse("input.txt")
	fmt.Printf("Task %d: %v\n", 1, task1And2(lists, 0))
	fmt.Printf("Task %d: %v\n", 2, task1And2(lists, 1))
}
