package main

import "fmt"

func testNil(va []string) {
	for _, v := range va {
		fmt.Println(v)
	}
}

func main() {
	testNil(nil)
}
