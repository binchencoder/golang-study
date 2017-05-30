package main

import "fmt"

func main() {
	ages := make(map[string]int) // create a map through make func, mapping from strings to ints
	ages["alice"] = 31
	ages["charlie"] = 34
	fmt.Printf("Create a map through make: %v\n", ages)

	ages1 := map[string]int{
		"chenbin":   28,
		"zhangxiao": 28,
	}
	fmt.Printf("Create a age map: %v\n", ages1)

	// Map中的元素通过key对应的下标语法访问
	key := "chenbin"
	fmt.Printf("Access age map: %v, through the key: %s, value: %d\n", ages1, key, ages1[key])

	// 使用内置的delete函数可以删除元素
	delete(ages, "alice")

	if age, ok := ages1[key]; ok {
		fmt.Printf("The key[%s] in the ages1 map, value is %d\n", key, age)
	}

	// 遍历map
	for name, age := range ages {
		fmt.Printf("Key: %s, Value: %d\n", name, age)
	}
}
