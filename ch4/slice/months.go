package main

import (
	"fmt"
)

var (
	months = [...]string{1: "January", 2: "February", 3: "March", 4: "April", 5: "May", 6: "June", 7: "July", 8: "August", 9: "September", 10: "October", 11: "November", 12: "December"}
)

func findSameMonth(Q2 []string, summer []string) {
	for _, s := range summer {
		for _, q := range Q2 {
			if s == q {
				fmt.Printf("%s appears in both\n", s)
			}
		}
	}
}

func main() {
	fmt.Printf("12 months of the year\n %v\n %#v\n", months, months)

	Q2 := months[4:7]
	fmt.Println(Q2)

	summer := months[6:9]
	fmt.Println(summer)

	findSameMonth(Q2, summer)
}
