package main

import (
	"testing"
)

var (
	g = parse("input.txt")
)

func BenchmarkTask1(b *testing.B) {
	b.StartTimer()
	task1(g, "XMAS")
}
func BenchmarkTask2(b *testing.B) {
	b.StartTimer()
	task2(g, "MAS")
}
