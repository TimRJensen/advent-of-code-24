package main

import (
	"testing"
)

var (
	lists = parse("input.txt")
)

func BenchmarkTask1(b *testing.B) {
	b.StartTimer()
	task1And2(lists, 0)
}
func BenchmarkTask2(b *testing.B) {
	b.StartTimer()
	task1And2(lists, 1)
}
