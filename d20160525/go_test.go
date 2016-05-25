package d20160525

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {
	slc := make([]int, 1, 3)
	fmt.Println(&(slc[0]))
	slc = append(slc, 1)
	fmt.Println(&(slc[0]))
	slc = append(slc, 2)
	fmt.Println(&(slc[0]))
	slc = append(slc, 3)
	fmt.Println(&(slc[0]))

	m := make(map[string]bool)
	fmt.Println(m)
	changeMap(m)
	fmt.Println(m)
}

func changeMap(m map[string]bool) {
	m["one"] = true
	m["two"] = false
}
