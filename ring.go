package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

func main() {
	// Manual to do list RIng
	// var data *ring.Ring = ring.New(5)

	// data.Value = "value 1"

	// data = data.Next()
	// data.Value = "value 2"

	// data = data.Next()
	// data.Value = "value 3"

	// data = data.Next()
	// data.Value = "value 4"

	// data = data.Next()
	// data.Value = "value 5"

	var data *ring.Ring = ring.New(5)
	for i := 0; i < data.Len(); i++ {
		data.Value = "Value " + strconv.FormatInt(int64(i), 10)
		data = data.Next()
	}

	data.Do(func(value any) {
		fmt.Println(value)
	})
}
