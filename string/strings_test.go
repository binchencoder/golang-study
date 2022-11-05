package string

import (
	"fmt"
	"strings"
	"testing"
)

func TestContains(t *testing.T) {
	contains := strings.Contains("entityType,entity", "e")
	fmt.Printf("Test strings.Contains %v\n", contains)
}

func TestSubStr(t *testing.T) {
	str := "1,2"
	idx := strings.LastIndex(str, ",")
	fmt.Printf("last '.'.index: %d \n", idx)

	var strArr []string
	if idx > -1 {
		str = str[:idx]
		strArr = strings.Split(str, ",")
	}

	fmt.Printf("str: %s, strArr: %v \n", str, strArr)
}

func StringsContains(array []string, val string) (index int) {
	index = -1
	for i := 0; i < len(array); i++ {
		if array[i] == val {
			index = i
			return
		}
	}
	return
}
