package main

import (
	"runtime"
	"strconv"
	"testing"
)

func AssertEqual(exp interface{}, act interface{}, t *testing.T) {
	_, file, line, _ := runtime.Caller(1)
	if exp != act {
		t.Error("expected ", exp, ", actual ", act, "\n", "line ", strconv.Itoa(line), " in file ", file)
	}
}

func TestAssert(t *testing.T) {
	AssertEqual("exp", "act", t)
	mp := make(map[string]interface{}, 10)
	t.Log(mp, len(mp))
}
