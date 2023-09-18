package modules

import (
	"fmt"
	"sync"
	"time"
)

func AsyncCallFunctions() {
	start := time.Now()

	channels := make(chan string, 128)
	waitGroups := &sync.WaitGroup{}

	go asyncFunc1(channels, waitGroups)
	waitGroups.Add(1)

	go asyncFunc2(channels, waitGroups)
	waitGroups.Add(1)

	go asyncFunc3(channels, waitGroups)
	waitGroups.Add(1)

	waitGroups.Wait()

	close(channels)

	fmt.Println("-------------- ASYNC CALL --------------")
	//we can for range on channels
	for channel := range channels {
		fmt.Println(channel)
	}

	fmt.Printf("*** Time Duration : %s *** \n", time.Since(start))
}

func asyncFunc1(chanel chan string, wg *sync.WaitGroup) {
	time.Sleep(170 * time.Millisecond)

	chanel <- "func 1 done"

	wg.Done()
}

func asyncFunc2(chanel chan string, wg *sync.WaitGroup) {
	time.Sleep(100 * time.Millisecond)

	chanel <- "func 2 done"

	wg.Done()
}

func asyncFunc3(chanel chan string, wg *sync.WaitGroup) {
	time.Sleep(300 * time.Millisecond)

	chanel <- "func 3 done"

	wg.Done()
}
