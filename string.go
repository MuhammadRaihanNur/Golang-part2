package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Hello, World!", "World"))
	fmt.Println(strings.Split("Hello, World!", " "))
	fmt.Println(strings.ToLower("Muhammad Raihan"))
	fmt.Println(strings.ToUpper("Raihan Muhammad"))
	fmt.Println(strings.Trim("adadad  Muhammad Raihan", "adadad  "))
	fmt.Println(strings.ReplaceAll("Muhammd Raihan Jelek Raihan Jelek", "Jelek", "Ganteng"))
}
