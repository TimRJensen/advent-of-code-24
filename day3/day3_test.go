package main

import (
	"testing"
)

var (
	buff = parse("input.txt")
)

func BenchmarkTask1(b *testing.B) {
	b.StartTimer()
	task1And2(buff, false)
}
func BenchmarkTask2(b *testing.B) {
	b.StartTimer()
	task1And2(buff, true)
}
