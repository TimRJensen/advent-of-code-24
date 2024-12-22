package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func task1(g *grid, txt string) int {
	res := 0
	ops := []func(int) (bool, int){
		g.left, g.right, g.up, g.down, g.upLeft, g.upRight, g.downLeft, g.downRight,
	}

	for i := 0; i < len(g.buff); i++ {
		if g.buff[i] != txt[0] {
			continue
		}

		for _, op := range ops {
			if g.isValidSequence(i, op, txt, 0) {
				res++
			}
		}
	}

	return res
}

func task2(g *grid, txt string) int {
	res := 0
	ops := []func(int) (bool, int){
		g.upLeft, g.upRight, g.downRight, g.downLeft,
	}

	for i := 0; i < len(g.buff); i++ {
		if g.buff[i] != txt[1] {
			continue
		}

		n := 0
		for j, op := range ops {
			ok, i := op(i)
			if !ok {
				continue
			}
			if g.isValidSequence(i, ops[(j+2)%len(ops)], txt, 0) {
				n += 1
			}
		}
		if n == 2 {
			res++
		}
	}

	return res
}

func parse(path string) *grid {
	buff, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	g := new(grid)
	for _, buff := range bytes.Split(buff, []byte{'\n'}) {
		g.buff = append(g.buff, bytes.TrimSpace(buff)...)
		if g.n == 0 {
			g.n = len(g.buff)
		}
	}
	return g
}

func main() {
	g := parse("input.txt")
	fmt.Printf("Task %v: %v\n", 1, task1(g, "XMAS"))
	fmt.Printf("Task %v: %v\n", 2, task2(g, "MAS"))
}
