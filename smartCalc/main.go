package main

import (
	"fmt"
)

func main() {
	arr := []int{0, 1, 2, 3, 4, 5, 3}
	fmt.Println(Calc(arr, 0, 1))
	fmt.Println(Calc(arr, 0, 5))
	fmt.Println(Calc(arr, 0, 0))
	fmt.Println(Calc(arr, 0, 6))
}

func Calc(arr []int, n1, n2 int) int {
	val := arr[n1 : n2+1]
	count := 0
	for _, v := range val {
		count += v
	}

	return count
}
