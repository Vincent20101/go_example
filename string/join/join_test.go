package main

import (
	"strings"
	"testing"
)

func BenchmarkPlus(b *testing.B) {
	str := "this is just a string"

	for i := 0; i < b.N; i++ {
		stringPlus(str)
	}
}

func BenchmarkSPrintf(b *testing.B) {
	str := "this is just a string"
	for i := 0; i < b.N; i++ {
		stringSprintf(str)
	}
}

func BenchmarkStringBuilder(b *testing.B) {
	str := "this is just a string"
	for i := 0; i < b.N; i++ {
		stringBuilder(str)
	}
}

func stringPlus(str string) string {
	s := ""
	for i := 0; i < 10000; i++ {
		s += str
	}
	return s
}

func stringSprintf(str string) string {
	s := ""
	for i := 0; i < 10000; i++ {
		s += str
	}
	return s
}

func stringBuilder(str string) string {
	builder := strings.Builder{}
	for i := 0; i < 100000; i++ {
		builder.WriteString(str)
	}
	return builder.String()
}

func BenchmarkStringBuilderPre(b *testing.B) {
	str := "this is just a string"
	for i := 0; i < b.N; i++ {
		stringBuilderPre(str)
	}
}

func stringBuilderPre(str string) string {
	builder := strings.Builder{}
	builder.Grow(1000000)
	for i := 0; i < 100000; i++ {
		builder.WriteString(str)
	}
	return builder.String()
}
