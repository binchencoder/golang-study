package arry

import (
	"fmt"
	"reflect"
	"testing"
)

func TestContains(t *testing.T) {
	contains := StringsContains([]string{"entityType", "entity"}, "entity")
	fmt.Printf("Test strings.Contains %d\n", contains)
}

func BenchmarkContains(b *testing.B) {
	sa := []string{"q", "w", "e", "r", "t"}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		Contains(sa, "r")
	}
}

func BenchmarkStringsContains(b *testing.B) {
	sa := []string{"q", "w", "e", "r", "t"}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		StringsContains(sa, "r")
	}
}

func Contains(array interface{}, val interface{}) (index int) {
	index = -1
	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		{
			s := reflect.ValueOf(array)
			for i := 0; i < s.Len(); i++ {
				if reflect.DeepEqual(val, s.Index(i).Interface()) {
					index = i
					return
				}
			}
		}
	}
	return
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
