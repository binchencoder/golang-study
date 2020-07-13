package main

import (
	"errors"
	"fmt"
)

func main() {
	err := err()
	if err != nil {
		return
	}

	defer func() {
		fmt.Println("end")
	}()
}

func err() error {
	return errors.New("Test error")
}
