package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regexp *regexp.Regexp = regexp.MustCompile(`e([a-z])o`)

	fmt.Println(regexp.MatchString("eko"))
	fmt.Println(regexp.MatchString("edo"))
	fmt.Println(regexp.MatchString("eLo"))

	fmt.Println(regexp.FindAllString("eko edo egi ego e1o eto eKo", 10))
}
