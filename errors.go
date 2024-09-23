package main

import (
	"errors"
	"fmt"
)

var (
	ValidationError = errors.New("validation error")
	NotFoundError   = errors.New("not found error")
)

func CerId(id string) error {
	if id == "" {
		return ValidationError
	}

	if id != "raihan" {
		return NotFoundError
	}

	return nil
}

func main() {
	err := CerId("")
	if err != nil {
		if errors.Is(err, ValidationError) {
			fmt.Println("Validation error")
		} else if errors.Is(err, NotFoundError) {
			fmt.Println("Not found error")
		} else {
			fmt.Println("Unknown error")
		}
	}

}
