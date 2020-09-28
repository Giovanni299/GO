package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println(Istwin("Hello", "world"))
	fmt.Println(Istwin("acb", "bca"))
	fmt.Println(Istwin("Lookout", "Outlook"))
}

func Istwin(a string, b string) bool {
	a = strings.ToLower(a)
	b = strings.ToLower(b)
	fmt.Println(a, b)

	a1 := strings.Split(a, "")
	sort.Strings(a1)
	b1 := strings.Split(b, "")
	sort.Strings(b1)
	fmt.Println(a1, b1)

	if strings.Join(a1, "") == strings.Join(b1, "") {
		return true
	}

	return false
}
