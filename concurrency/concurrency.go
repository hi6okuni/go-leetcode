package concurrency

import (
	"fmt"
	"sync"
)

func Test() {
	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	fmt.Println("1st goroutine sleeping...")
	// 	time.Sleep(1 * time.Second)
	// }()

	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	fmt.Println("2nd goroutine sleeping...")
	// 	time.Sleep(2 * time.Second)
	// }()

	// wg.Wait()
	// fmt.Println("All goroutines complete.")

	// hello := func(wg *sync.WaitGroup, id int) {
	// 	defer wg.Done()
	// 	fmt.Printf("Hello from %v!\n", id)
	// }

	// const numGreeters = 5
	// var wg sync.WaitGroup
	// wg.Add(numGreeters)
	// for i := 0; i < numGreeters; i++ {
	// 	go hello(&wg, i+1)
	// }
	// wg.Wait()

	var count int
	var lock sync.Mutex

	increament := func() {
		lock.Lock()
		defer lock.Unlock()
		count++
		fmt.Printf("Incrementing: %d\n", count)
	}

	decreament := func() {
		lock.Lock()
		defer lock.Unlock()
		count--
		fmt.Printf("Decrementing: %d\n", count)
	}

	var arithmetic sync.WaitGroup
	for i := 0; i <= 5; i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			increament()
		}()
	}

	for i := 0; i <= 5; i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			decreament()
		}()
	}

	arithmetic.Wait()

	fmt.Println("Arithmetic complete.")
}
