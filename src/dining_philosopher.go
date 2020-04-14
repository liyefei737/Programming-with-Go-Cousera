package main

import (
	"fmt"
	"sync"
)

type ChopStick struct{ sync.Mutex }

type Philosopher struct {
	leftCS  *ChopStick
	rightCS *ChopStick
	id      int
}

func (p Philosopher) eat(eatChan chan int, wg *sync.WaitGroup) {
	for i := 0; i < 3; i++ {
		eatChan <- p.id
		p.leftCS.Lock()
		p.rightCS.Lock()

		fmt.Println("Starting to eat", p.id)
		fmt.Println("Finishing eat", p.id)

		p.rightCS.Unlock()
		p.leftCS.Unlock()
		<-eatChan
	}
	wg.Done()
}

func main() {

	chopsticks := make([]*ChopStick, 5)
	philosophers := make([]*Philosopher, 5)
	var wg sync.WaitGroup

	// channel of capacity 2
	var eatChan = make(chan int, 2)

	for i := 0; i < 5; i++ {
		chopsticks[i] = new(ChopStick)
	}

	for i := 0; i < 5; i++ {
		philosophers[i] = &Philosopher{chopsticks[i], chopsticks[(i+1)%5], i + 1}
	}

	wg.Add(5)

	// each philospher tries to eat
	for i := 0; i < 5; i++ {
		go philosophers[i].eat(eatChan, &wg)
	}

	wg.Wait()

}
