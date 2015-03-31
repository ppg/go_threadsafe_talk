package main

import (
	"fmt"
	//"runtime"
	"sync"
	"testing"
	"time"
)

const loop = 100
const delay = time.Millisecond

func BenchmarkFmtPrint(b *testing.B) {
	wg := &sync.WaitGroup{}
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func() {
			for i := 0; i < loop; i++ {
				fmt.Print("")
				time.Sleep(delay)
			}
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
			for i := 0; i < loop; i++ {
				mutex.Lock()
				fmt.Print("")
				mutex.Unlock()
				time.Sleep(delay)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func BenchmarkChannelFmtPrint(b *testing.B) {
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
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func() {
			for i := 0; i < loop; i++ {
				c <- ""
				time.Sleep(delay)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	q <- struct{}{}
}
