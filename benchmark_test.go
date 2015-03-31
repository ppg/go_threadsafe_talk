package main

import (
	"fmt"
	"sync"
	"testing"
)

// Runs f in n go routines, waiting until they all complete
func benchmarkNGoroutines(n int, f func()) {
	wg := &sync.WaitGroup{}
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			f()
		}()
	}
	wg.Wait()
}

func benchmarkFmtPrint(routines int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchmarkNGoroutines(routines, func() {
			fmt.Print("")
		})
	}
}

func BenchmarkFmtPrint1(b *testing.B)    { benchmarkFmtPrint(1, b) }
func BenchmarkFmtPrint10(b *testing.B)   { benchmarkFmtPrint(10, b) }
func BenchmarkFmtPrint100(b *testing.B)  { benchmarkFmtPrint(100, b) }
func BenchmarkFmtPrint1000(b *testing.B) { benchmarkFmtPrint(1000, b) }

func benchmarkMutexFmtPrint(routines int, b *testing.B) {
	var mutex sync.Mutex
	for i := 0; i < b.N; i++ {
		benchmarkNGoroutines(routines, func() {
			mutex.Lock()
			fmt.Print("")
			mutex.Unlock()
		})
	}
}

func BenchmarkMutexFmtPrint1(b *testing.B)    { benchmarkMutexFmtPrint(1, b) }
func BenchmarkMutexFmtPrint10(b *testing.B)   { benchmarkMutexFmtPrint(10, b) }
func BenchmarkMutexFmtPrint100(b *testing.B)  { benchmarkMutexFmtPrint(100, b) }
func BenchmarkMutexFmtPrint1000(b *testing.B) { benchmarkMutexFmtPrint(1000, b) }

func benchmarkChannelFmtPrint(routines int, b *testing.B) {
	c := make(chan string)
	q := make(chan struct{})
	go func() {
		for {
			select {
			case s := <-c:
				fmt.Print(s)
			case <-q:
				return
			}
		}
	}()

	for i := 0; i < b.N; i++ {
		benchmarkNGoroutines(routines, func() {
			c <- ""
		})
	}
	q <- struct{}{}
}

func BenchmarkChannelFmtPrint1(b *testing.B)    { benchmarkChannelFmtPrint(1, b) }
func BenchmarkChannelFmtPrint10(b *testing.B)   { benchmarkChannelFmtPrint(10, b) }
func BenchmarkChannelFmtPrint100(b *testing.B)  { benchmarkChannelFmtPrint(100, b) }
func BenchmarkChannelFmtPrint1000(b *testing.B) { benchmarkChannelFmtPrint(1000, b) }
