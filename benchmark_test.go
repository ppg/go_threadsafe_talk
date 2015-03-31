package main

import (
	"fmt"
	"sync"
	"testing"
)

func benchmarkFmtPrint(routines int, b *testing.B) {
	for j := 0; j < b.N; j++ {
		wg := &sync.WaitGroup{}
		for i := 0; i < routines; i++ {
			wg.Add(1)
			go func() {
				fmt.Print("")
				wg.Done()
			}()
		}
		wg.Wait()
	}
}
func BenchmarkFmtPrint1(b *testing.B)    { benchmarkFmtPrint(1, b) }
func BenchmarkFmtPrint10(b *testing.B)   { benchmarkFmtPrint(10, b) }
func BenchmarkFmtPrint100(b *testing.B)  { benchmarkFmtPrint(100, b) }
func BenchmarkFmtPrint1000(b *testing.B) { benchmarkFmtPrint(1000, b) }

func benchmarkMutexFmtPrint(routines int, b *testing.B) {
	for j := 0; j < b.N; j++ {
		wg := &sync.WaitGroup{}
		var mutex sync.Mutex
		for i := 0; i < routines; i++ {
			wg.Add(1)
			go func() {
				mutex.Lock()
				fmt.Print("")
				mutex.Unlock()
				wg.Done()
			}()
		}
		wg.Wait()
	}
}
func BenchmarkMutexFmtPrint1(b *testing.B)    { benchmarkMutexFmtPrint(1, b) }
func BenchmarkMutexFmtPrint10(b *testing.B)   { benchmarkMutexFmtPrint(10, b) }
func BenchmarkMutexFmtPrint100(b *testing.B)  { benchmarkMutexFmtPrint(100, b) }
func BenchmarkMutexFmtPrint1000(b *testing.B) { benchmarkMutexFmtPrint(1000, b) }

func benchmarkChannelFmtPrint(routines int, b *testing.B) {
	for j := 0; j < b.N; j++ {
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

		wg := &sync.WaitGroup{}
		for i := 0; i < routines; i++ {
			wg.Add(1)
			go func() {
				c <- ""
				wg.Done()
			}()
		}
		wg.Wait()
		q <- struct{}{}
	}
}

func BenchmarkChannelFmtPrint1(b *testing.B)    { benchmarkChannelFmtPrint(1, b) }
func BenchmarkChannelFmtPrint10(b *testing.B)   { benchmarkChannelFmtPrint(10, b) }
func BenchmarkChannelFmtPrint100(b *testing.B)  { benchmarkChannelFmtPrint(100, b) }
func BenchmarkChannelFmtPrint1000(b *testing.B) { benchmarkChannelFmtPrint(1000, b) }
