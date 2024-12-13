package main

import (
	"testing"
)

var (
	n, left, right = parse("input.txt")
)

func BenchmarkTask1(b *testing.B) {
	b.StartTimer()
	task1(left, right)
}
func BenchmarkTask2(b *testing.B) {
	b.StartTimer()
	task2(n, left, right)
}
