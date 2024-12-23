package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"slices"

	"github.com/TimRJensen/aoc2024/util"
)

func partition(seq []int, rule [][]int, low, high int) int {
	pivot := seq[low]
	i := low - 1
	j := high + 1

	for {
		// Any seq[i] without pivot in rule[seq[i]] succeeds the pivot
		for {
			i++
			if _, ok := slices.BinarySearch(rule[seq[i]], pivot); !ok || seq[i] == pivot {
				break
			}
		}
		// Any seq[j] with pivot in rules[seq[j]] precedes the pivot
		for {
			j--
			if _, ok := slices.BinarySearch(rule[seq[j]], pivot); ok || seq[j] == pivot {
				break
			}
		}

		if i >= j {
			return j
		}
		seq[i], seq[j] = seq[j], seq[i]
	}
}

func quicksort(seq []int, rule [][]int, low, high int) {
	if low < high {
		p := partition(seq, rule, low, high)
		quicksort(seq, rule, low, p)
		quicksort(seq, rule, p+1, high)
	}
}

func task1And2(seqs [][]int, rules [][]int, sort bool) int {
	sum := 0
	for _, seq := range seqs {
		i := len(seq) - 2
		for ; i >= 0; i-- {
			if _, ok := slices.BinarySearch(rules[seq[i]], seq[i+1]); !ok {
				break
			}
		}

		if i == -1 && !sort {
			sum += seq[len(seq)/2]
		}
		if i > -1 && sort {
			quicksort(seq, rules, 0, len(seq)-1)
			sum += seq[len(seq)/2]
		}
	}
	return sum
}

func parse(path string) ([][]int, [][]int) {
	buff, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	rules := make([][]int, 124) // Assuming a input number < 124. Should probably be generalized to a map[int][]int
	seqs := make([][]int, 0, 124)
	for _, buff := range bytes.Split(buff, []byte{'\n'}) {
		if buff = bytes.TrimSpace(buff); len(buff) == 0 {
			continue
		}

		fields := bytes.Split(buff, []byte{'|'})
		if len(fields) == 2 {
			idx, v := util.Atoi(fields[0]), util.Atoi(fields[1])
			rules[idx] = append(rules[idx], v)
			continue
		}

		seqs = append(seqs, make([]int, 0, 16))
		for _, field := range bytes.Split(fields[0], []byte{','}) {
			seqs[len(seqs)-1] = append(seqs[len(seqs)-1], util.Atoi(field))
		}
	}

	for _, rule := range rules {
		slices.Sort(rule)
	}

	return seqs, rules
}

func main() {
	seqs, rules := parse("input.txt")
	fmt.Printf("Task %v: %v\n", 1, task1And2(seqs, rules, false))
	fmt.Printf("Task %v: %v\n", 2, task1And2(seqs, rules, true))
}
