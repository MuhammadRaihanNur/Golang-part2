package main

import (
	"fmt"
	"strconv"
)

func main() {
	result, err := strconv.ParseBool("false")
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(result)
	}

	resultInt, err := strconv.Atoi("706")
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(resultInt)
	}

	angka := strconv.FormatInt(888, 3)
	fmt.Println(angka)

	stringInt := strconv.Itoa(666)
	fmt.Println(stringInt)

}
