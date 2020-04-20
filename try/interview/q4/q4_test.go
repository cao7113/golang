package q4

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

func TestA(t *testing.T) {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("A: ", i)
			wg.Done()
		}()
	}
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("B: ", i)

			wg.Done()
		}(i)
	}
	wg.Wait()
}
