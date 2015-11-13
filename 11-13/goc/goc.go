package main

/*
#include "hello.h"
#include <stdio.h>

void sayHi() {
  printf("Hi");
  sayHello();
}
*/
import "C"

func main() {
	C.sayHi()
}
