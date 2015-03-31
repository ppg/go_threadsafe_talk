package main

import (
	"fmt"
	"sync"
	"testing"
)

func BenchmarkFmtPrint(b *testing.B) {
	wg := &sync.WaitGroup{}
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func() {
			fmt.Print("")
			wg.Done()
		}()
	}
	wg.Wait()
}

func BenchmarkMutexFmtPrint(b *testing.B) {
	wg := &sync.WaitGroup{}
	var mutex sync.Mutex
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func() {
			mutex.Lock()
			defer mutex.Unlock()
			fmt.Print("")
			wg.Done()
		}()
	}
	wg.Wait()
}

func BenchmarkChannelFmtPrint(b *testing.B) {
	c := make(chan string)
	q := make(chan struct{})
	go func() {
		select {
		case s := <-c:
			fmt.Print(s)
		case <-q:
		}
	}()

	wg := &sync.WaitGroup{}
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func() {
			c <- ""
			wg.Done()
		}()
	}
	wg.Wait()
	q <- struct{}{}
}
