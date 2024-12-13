package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func task1(left, right []int) int {
	slices.Sort(left)
	slices.Sort(right)

	sum := 0
	for i := 0; i < len(left); i++ {
		if l, r := left[i], right[i]; l < r {
			sum += r - l
		} else {
			sum += l - r
		}
	}
	return sum
}

func task2(n int, left, right []int) int {
	counts := make([]int, n+1)
	for i := 0; i < len(right); i++ {
		counts[right[i]]++
	}

	sum := 0
	for i := 0; i < len(left); i++ {
		sum += left[i] * counts[left[i]]
	}
	return sum
}

func parse(path string) (int, []int, []int) {
	buff, err := os.ReadFile(path)
	if err != nil {
		log.Fatalln(err)
	}

	ns := make([][]int, 2, 32)
	n := -1
	for _, line := range strings.Split(string(buff), "\n") {
		for j, field := range strings.Fields(line) {
			if v, err := strconv.Atoi(field); err == nil {
				ns[j] = append(ns[j], v)
				if v > n {
					n = v
				}
			}
		}
	}

	return n, ns[0], ns[1]
}

func main() {
	n, left, right := parse("input.txt")
	fmt.Printf("Task %d: %v\n", 1, task1(left, right))
	fmt.Printf("Task %d: %v\n", 2, task2(n, left, right))
}
