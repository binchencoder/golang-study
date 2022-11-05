package date

import (
	"fmt"
	"testing"
	"time"
)

func TestDateParse(t *testing.T) {
	gmtCreate, _ := time.Parse("2006-01-02", "2022-07-25T18:00:43.029Z")
	fmt.Printf("Test time.Parse %s\n", gmtCreate)
}
