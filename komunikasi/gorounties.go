package main

import (
	"fmt"
	"time"
)

func printPesan(msg string) {
	for i := 0; i < 5; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	go printPesan("Tugas 1")
	go printPesan("Tugas 2")

	time.Sleep(time.Second * 5)
	fmt.Println("semua tugas selesai")
}
