package main

import (
	"container/list"
	"fmt"
)

func main() {
	var data *list.List = list.New()

	data.PushBack("Raihan")
	data.PushBack("Nur")
	data.PushBack("Ahsan")

	var head *list.Element = data.Front()
	fmt.Println(head.Value)

	lanjut := head.Next()
	fmt.Println(lanjut.Value)

	lanjut = lanjut.Next()
	fmt.Println(lanjut.Value)

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
