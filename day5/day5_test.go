package main

import (
	"testing"
)

var (
	seqs, rules = parse("input.txt")
)

func BenchmarkTask1(b *testing.B) {
	b.StartTimer()
	task1And2(seqs, rules, false)
}
func BenchmarkTask2(b *testing.B) {
	b.StartTimer()
	task1And2(seqs, rules, true)
}
