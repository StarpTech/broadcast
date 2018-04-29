package main

import (
	"fmt"
	"sync"

	"github.com/starptech/broadcast"
)

func main() {
	s := new(sync.WaitGroup)
	b := make(broadcast.Broadcast) // Broadcast channels are constructed exactly like regular channels
	s.Add(2)
	go func() {
		fmt.Println(b.Receive())
		s.Done()
	}()
	go func() {
		fmt.Println(b.Receive())
		s.Done()
	}()
	b <- 1
	s.Wait()
}
