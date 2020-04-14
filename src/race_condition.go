package main

import (
	"fmt"
	"time"
)

/*
This program has a race condition such that it's possible that it will print any value among 0, 1,2, and 3, 4.
Bacause after the short delay, we do not know which goroutines have already executed.
*/

func main() {
	var data int
	go func() {
		data++
	}()
	go func() {
		data++
	}()
	go func() {
		data++
	}()
	go func() {
		data++
	}()
	time.Sleep(8 * time.Microsecond)

	fmt.Printf("the value is %v.\n", data)

}
