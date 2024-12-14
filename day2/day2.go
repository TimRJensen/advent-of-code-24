package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func task1(lists [][]int) int {
	// Case 1: All ascending 1..9 i==j
	// Case 2: All descending 9..1 == ascending reverse(9..1)->1..9
	sum := 0
	for _, lst := range lists {
		i, j := 0, len(lst)-1
		for i < len(lst)-1 && j > 0 {
			if lst[i+1]-lst[i] > 0 && lst[i+1]-lst[i] < 4 {
				i++
				continue
			}
			if lst[j-1]-lst[j] > 0 && lst[j-1]-lst[j] < 4 {
				j--
				continue
			}
			break
		}
		if i == len(lst)-1 || j == 0 {
			sum++
		}
	}

	return sum
}

func task2(lists [][]int) int {
	sum := 0
	for _, lst := range lists {
		i, j, delta := 0, len(lst)-1, 0
		for i < len(lst)-1 && j > 0 {
			if lst[i+1]-lst[i] > 0 && lst[i+1]-lst[i] < 4 {
				i++
				continue
			}
			if lst[j-1]-lst[j] > 0 && lst[j-1]-lst[j] < 4 {
				j--
				continue
			}
			if delta > 0 {
				break
			}
			i++
			j--
			delta++
		}
		if delta < 2 && (i == len(lst)-1 || j == 0) {
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
	fmt.Printf("Task %d: %v\n", 1, task1(lists))
	fmt.Printf("Task %d: %v\n", 2, task2(lists))
}
