package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

const (
	xmas = "XMAS"
	mas  = "MAS"
)

var (
	table = [26]int{}
)

type grid struct {
	buff []byte
	n    int
}

func (g grid) left(i int) (bool, int) {
	if i < 0 || i >= len(g.buff) || (i-1)/g.n < i/g.n {
		return false, i
	}
	return true, i - 1
}

func (g grid) right(i int) (bool, int) {
	if i < 0 || i >= len(g.buff) || (i+1)/g.n > i/g.n {
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

func (g grid) isValidSequence(i int, op func(int) (bool, int), txt string) bool {
	if ok, j := op(i); !ok {
		return false
	} else {
		if txt[table[g.buff[i]-'A']+1] != txt[table[g.buff[j]-'A']] {
			return false
		}
		if txt[table[g.buff[j]-'A']] == txt[len(txt)-1] {
			return true
		}
		return g.isValidSequence(j, op, txt)
	}
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
		if g.buff[i] != xmas[0] {
			continue
		}

		for _, op := range ops {
			if ok := g.isValidSequence(i, op, xmas); ok {
				res++
			}
		}
	}

	return res
}

func task2(g *grid) int {
	res := 0
	ops := []func(int) (bool, int){
		g.upLeft, g.downRight, g.downLeft, g.upRight,
	}
	//TODO: update the table and all is well.

	for i := 0; i < len(g.buff); i++ {
		if g.buff[i] != mas[1] {
			continue
		}

		flag := true
		for j := 0; j < len(ops); j += 2 {
			if ok, k := ops[j](i); !ok {
				continue
			} else {
				flag = flag && g.isValidSequence(k, ops[j+1], mas)
			}
		}

		if flag {
			res += 1
		}
	}

	return res
}

func main() {
	g := parse("input.txt")
	fmt.Printf("Task %v: %v\n", 1, task1(g))
	fmt.Printf("Task %v: %v\n", 2, task2(g))
}

func init() {
	table[23] = 0
	table[12] = 1
	table[0] = 2
	table[18] = 3
}
