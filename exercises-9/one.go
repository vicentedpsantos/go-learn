package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(2)

	for i := 0; i < 2; i++ {
		go func() {
			time.Sleep(time.Second)

			wg.Done()
		}()

		fmt.Println("From within go routine", i)
	}

}
