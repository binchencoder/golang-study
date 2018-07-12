package main

import (
	"fmt"
	"time"
)

// Employee of test struct
type Employee struct {
	ID        int
	Name      string
	Address   string
	Dob       time.Time
	Position  string
	Salary    int
	ManagerID int
}

// AwardAnnualRaise of test
func AwardAnnualRaise(e *Employee) {
	e.Salary = e.Salary * 105
}

func main() {
	e := &Employee{Salary: 3}
	fmt.Printf("Before: %+v\n", e)

	AwardAnnualRaise(e)

	fmt.Printf("After: %+v\n", e)
}
