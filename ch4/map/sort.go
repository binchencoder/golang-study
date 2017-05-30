package main

import (
	"fmt"
	"sort"
)

func main() {
	ages := map[string]int{
		"zhangs": 31,
		"wangw":  33,
		"lis":    29,
	}

	names := make([]string, 0, len(ages))
	for name := range ages {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}
}
