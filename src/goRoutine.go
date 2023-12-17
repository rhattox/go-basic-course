package main

import (
	"fmt"
	"sync"
)

func printNumbers(wg *sync.WaitGroup) {
	// Ensure that the WaitGroup is decremented at the end
	defer wg.Done()

	for i := 5; i <= 10; i++ {
		fmt.Printf("%d ", i)
	}

	fmt.Println(" - Goroutine finished")
}

func printStrings(wg *sync.WaitGroup){
	defer wg.Done()

	stringArray := [...]string{ "1", "2", "blue", "yellow", "card"}
	for _, value := range stringArray {
		fmt.Printf("%s ",value)
	}
	fmt.Println(" - Goroutine finished")
}

func main() {
	// Create a WaitGroup to wait for goroutines to finish
	var wg sync.WaitGroup

	// Increment the WaitGroup counter
	wg.Add(3)

	// Launch two goroutines
	go printNumbers(&wg)
	go printNumbers(&wg)
	go printStrings(&wg)

	// Wait for all goroutines to finish
	wg.Wait()

	fmt.Println("Main function exiting")
}
