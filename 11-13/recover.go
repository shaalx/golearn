package main

import (
	"fmt"
	// "sync"
	// "time"
)

func main() {
	// c := make(chan bool, 5)
	// go func() {
	// 	fmt.Println("go read:", <-c)
	// }()
	// c <- true
	// c <- false
	// c <- true
	// c <- true
	// // c <- true
	// close(c)
	// fmt.Println(c)

	// // for i := range c {
	// // 	fmt.Println(i)
	// // }
	// fmt.Println(len(c))

	// cc := make(chan bool, 1)
	// cc <- true
	// DO(cc)
	// time.Sleep(1e9)
	// <-make(chan bool)
	defer func() {
		if err := recover(); nil != err {
			fmt.Println(err)
			if err == "panic2" {
				fmt.Println("panic is panic2")
			}

			fmt.Println("\n-----------------end in recover")
		}
	}()

	Do3()
	Do2()
	fmt.Println("\n-----------------end")
}

func Do2() {
	panic("panic2")
}

func Do3() {
	panic("panic3")
}

// func DO(c chan bool) {
// 	panic("panic1")
// 	an := make(chan bool, 1)
// 	an <- true
// 	select {
// 	case <-c:
// 		println("read from c")
// 	case <-an:
// 		println("阻塞了")

// 	}
// 	// l := sync.Mutex{}
// 	// l.Lock()
// }
