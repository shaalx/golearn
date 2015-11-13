package main

import (
	"fmt"
	// "os"
	// "io"
	// "unsafe"
)

// System.out.println(a+b)
func main() {
	i := 10
	j := (byte)(i)
	fmt.Printf("%x", j)
	// var a [3]int
	// a = [3]int{1, 2, 3}
	// for _, it := range "我是" {
	// 	fmt.Println(string(it))
	// }

	var ii complex64
	fmt.Println(ii, real(ii), imag(ii))

	var b []byte
	b = []byte{1, 2, 3}
	change(b)
	fmt.Println(b)

	var m map[string]User
	m = make(map[string]User)
	m["a"] = User{Name: "xiaoming"}

	fmt.Println(m)
	for k, v := range m {
		fmt.Println(k, v)
	}
	bb, ok := m["b"]
	if ok {
		fmt.Println(bb)
	} else {
		fmt.Println("nil--", bb)
	}
	if nil == b {

	}
	// delete(m, "a")

	fmt.Println(m, len(m))
	s := make([]int, 5, 10)
	fmt.Println(len(s), cap(s))
	s = []int{1, 2, 3, 4, 5}
	s = append(s, 6)
	// s = []int{1, 2, 3, 4, 5, 6}
	fmt.Println(len(s), cap(s))

	var c chan int
	c = make(chan int, 2)
	// c <- 10
	wc := (chan<- int)(c)
	wc <- 11
	rc := (<-chan int)(c)
	content := <-rc
	fmt.Println(content)
	var fc func([]byte) bool
	fc = change
	fc(b)

	var iu IUser
	iu = &User{Name: "xiaoming"}
	iu = &UserB{Name: "xiaoming_b"}
	iu.Change("lili")
	fmt.Println(iu)
}

func change(b []byte) bool {
	b[0] = 100
	return true
}

type UserB struct {
	Name string
	Age  int
}

func (u *UserB) Change(n string) (bool, int) {
	u.Name = n
	return true, 0
}

type User struct {
	Name string
	Age  int
}

func (u *User) Change(n string) (bool, int) {
	u.Name = n
	return true, 1
}

type IUser interface {
	Change(n string) (bool, int)
}

func A() {
	var a interface{}
	a = User{Name: "jj"}

}

func func_name(a ...interface{}) {
	l := len(a)
	switch a.(type) {
	}
	func_name("sss", "sss")
	s()
	err := recover()
	if err != nil {

	}
}

func s() {
	panic("ssss")
	slice := make([]int, 9, 90)
	m := make(map[string]interface{})
	c := make(chan bool, 10)
	for i := range c {

	}
	close(c)
	for {
		select {
		case i := <-c:
			fmt.Println(i)
		case c <- true:
			fmt.Println("in")
		}
	}
}
