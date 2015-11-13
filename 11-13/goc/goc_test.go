package goc

/*
#include <hello.h>
*/
import "C"

func ExampleHello() {
	t.Log("CGO:")
	C.sayHello()
	// Output: CGO:
	// Hello
	//
}
