package main

import (
	"time"
)

func main() {
	c := make(chan bool)
	go func() {
		for {
			time.Sleep(1e9)
			c <- true
			// c = nil
			close(c)
			break
		}
	}()
	for {
		// cin, ok := <-c
		// println(cin, ok)
		select {
		case cin, ok := <-c:
			println(cin, ok)
			// 	// default:
			// 	// println("nothing")
		}
	}
}
