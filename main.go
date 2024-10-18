package main

import (
	"fmt"
	"sync"
)

type in struct {
	value    int32
	oddChan  chan int32
	evenChan chan int32
}

var serverChan chan in

func Server() {
	for v := range serverChan {
		if v.value%2 == 0 {
			v.evenChan <- v.value
		} else {
			v.oddChan <- v.value
		}
	}
}

func main() {
	serverChan = make(chan in)
	list := []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	oddChan := make(chan int32)
	evenChan := make(chan int32)

	odds, evens := []int32{}, []int32{}

	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		for v := range oddChan {
			odds = append(odds, v)
		}
		wg.Done()
	}()

	go func() {
		for v := range evenChan {
			evens = append(evens, v)
		}
		wg.Done()
	}()

	go Server()
	go func() {
		for _, v := range list {
			serverChan <- in{v, oddChan, evenChan}

		}
		close(oddChan)
		close(evenChan)
	}()
	wg.Wait()

	// wait until all the task in the waiting group are done
	fmt.Println("odds")
	for _, result := range odds {
		fmt.Printf("%d\n", result)
	}
	fmt.Println("evens")
	for _, result := range evens {
		fmt.Printf("%d\n", result)
	}

}
