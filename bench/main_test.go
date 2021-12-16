package main

import "testing"

func BenchmarkConcatStr1(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ConcatStr1("yuan", "zhen")
	}
}

func BenchmarkConcatStr2(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ConcatStr2("yuan", "zhen")
	}
}
func BenchmarkConcatStr3(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ConcatStr3("yuan", "zhen")
	}
}
func BenchmarkConcatStr4(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ConcatStr4("yuan", "zhen")
	}
}
