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
