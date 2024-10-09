package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

type Player []User

func (player Player) Len() int {
	return len(player)
}

func (player Player) Less(i, j int) bool {
	return player[i].Age < player[j].Age
}

func (player Player) Swap(i, j int) {
	player[i], player[j] = player[j], player[i]
}

func main() {
	users := []User{
		{"Muhammad", 20},
		{"Ahmad", 25},
		{"Ali", 18},
		{"Diah", 34},
	}

	sort.Sort(Player(users))

	fmt.Println(users)
}
