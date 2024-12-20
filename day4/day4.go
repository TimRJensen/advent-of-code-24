package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

const (
	txt = "XMAS"
)

var (
	table = [26]int{}
)

type grid struct {
	buff []byte
	n    int
}

func (g grid) left(i int) (bool, int) {
	if i < 0 || (i-1)/g.n < i/g.n {
		return false, i
	}
	return true, i - 1
}

func (g grid) right(i int) (bool, int) {
	if i < 0 || (i+1)/g.n > i/g.n {
		return false, i
	}
	return true, i + 1
}

func (g grid) up(i int) (bool, int) {
	if i/g.n == 0 {
		return false, i
	}
	return true, i - g.n
}

func (g grid) down(i int) (bool, int) {
	if i/g.n == g.n-1 {
		return false, i
	}
	return true, i + g.n
}

func (g grid) upLeft(i int) (bool, int) {
	return g.left(i - g.n)
}

func (g grid) upRight(i int) (bool, int) {
	return g.right(i - g.n)
}

func (g grid) downLeft(i int) (bool, int) {
	return g.left(i + g.n)
}

func (g grid) downRight(i int) (bool, int) {
	return g.right(i + g.n)
}

func (g grid) isValidSequence(i int, op func(int) (bool, int)) bool {
	if ok, j := op(i); ok {
		if txt[table[g.buff[j]-'A']] == txt[len(txt)-1] {
			return true
		}
		return g.isValidSequence(j, op)
	}
	return false
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

func task1(g *grid) int {
	res := 0
	ops := []func(int) (bool, int){
		g.left, g.right, g.up, g.down, g.upLeft, g.upRight, g.downLeft, g.downRight,
	}

	for i := 0; i < len(g.buff); i++ {
		if g.buff[i] != txt[0] {
			continue
		}

		for _, op := range ops {
			if ok := g.isValidSequence(i, op); ok {
				res++
			}
		}
	}

	return res
}

func main() {
	g := parse("input.txt")
	fmt.Printf("Task %v: %v\n", 1, task1(g))
}

func init() {
	table[23] = 0
	table[12] = 1
	table[0] = 2
	table[18] = 3
}
