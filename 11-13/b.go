package maint

func main2() {
	str := "Hello,中国"
	for i, it := range str {
		println(i, it, string(it))
	}
}

func Tt() int {
	println("ok")
	return 10
}
