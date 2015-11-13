package maint

import "testing"

func TestT(t *testing.T) {
	t.Log("hello, this is a test")
	ret := Tt()
	if ret == 10 {
		t.Log(ret)
	} else {
		t.Error("not 10.")
	}
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// b.Log(i)
	}
}

func TestChan(t *testing.T) {
	c := make(chan bool, 5)
	c <- true
	c <- true
	c <- true
	c <- true
	c <- true
	t.Log(c)
	n := 0
	for i := range c {
		n++
		t.Log(n)
		t.Log(i)
	}
	// ii, ok := <-c
	// t.Log(ii, ok)
}
