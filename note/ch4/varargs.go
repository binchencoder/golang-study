package main

import "fmt"

func test(s string, a ...int) {
	fmt.Printf("%T, %v\n", a, a)
}

func test1(a ...int) {
	fmt.Println(a)
}

func main() {
	test("abc")

	a := [3]int{10, 20, 30}
	test1(a[:]...)
}