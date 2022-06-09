package main

import "testing"

func BenchmarkConcatStr1(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TimerFunc()
	}
}
