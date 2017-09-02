package main

import ("testing"
"strconv")
// "fmt")

func bFib(n int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		input := strconv.Itoa(n)
		fib_processer(input)
	}
}

func BenchmarkFib1(b *testing.B) { bFib(10,b) }
func BenchmarkFib2(b *testing.B) { bFib(100,b) }
func BenchmarkFib3(b *testing.B) { bFib(1000,b) }
func BenchmarkFib4(b *testing.B) { bFib(10000,b) }
