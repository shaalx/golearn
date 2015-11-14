package main

import (
	"log"
	"time"
)

var a, b, c int

func main() {
	go func() {
		a = 1
		b = 2
	}()
	go func() {
		c = a + 2
	}()
	log.Println(a, b, c)
	time.Sleep(1e9)
}

type Ibase interface {
	Do()
	Doo()
}
