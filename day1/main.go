package main

import (
	"bufio"
	"log"
	"os"
	"sync"
)

func main() {
	s := bufio.NewScanner(os.Stdin)

	c := make(chan int)
	var wg sync.WaitGroup
	for s.Scan() {
		wg.Add(1)
		go func(str string) {
			c <- Solve(str)
			wg.Done()
		}(s.Text())
	}

	go func() {
		wg.Wait()
		close(c)
	}()

	res := 0
	for i := range c {
		res += i
	}

	log.Println("Result:", res)
}
