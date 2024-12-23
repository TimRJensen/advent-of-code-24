package main

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"github.com/TimRJensen/aoc2024/util"
)

func task1And2(buff []byte, useDo bool) int {
	sum, flag, l := 0, true, newLexer(buff)
	for tok := l.next(); tok.typ != EOF; tok = l.next() {
		if tok.typ == BAD_FUNC {
			continue
		}

		switch string(tok.val) {
		case Do:
			flag = true
		case Dont:
			flag = false
		default:
			if useDo && !flag {
				continue
			}

			n := 1
			for _, buff := range bytes.Split(tok.val, []byte{','}) {
				n *= util.Atoi(buff)
			}
			sum += n
		}

	}
	return sum
}

func parse(path string) []byte {
	buff, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return bytes.TrimSpace(buff)
}

func main() {
	buff := parse("input.txt")
	fmt.Printf("Task %v: %v\n", 1, task1And2(buff, false))
	fmt.Printf("Task %v: %v\n", 2, task1And2(buff, true))
}
