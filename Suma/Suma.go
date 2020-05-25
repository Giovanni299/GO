package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("=================SUMA==================")

	for d := 0; d < 10; d++ {
		//fmt.Println("D: " + strconv.Itoa(d))
		for o := 0; o < 10; o++ {
			//fmt.Println("O: " + strconv.Itoa(o))
			for s := 0; s < 10; s++ {
				//fmt.Println("S: " + strconv.Itoa(s))
				dos := s + (o * 10) + (d * 100)
				ocho := dos * 4
				cal := s * 4
				if cal >= 10 {
					cal = cal - (cal/10)*10
				}

				if cal == o && ocho/1000 == o {
					fmt.Println("DOS: " + strconv.Itoa(dos) + "     D: " + strconv.Itoa(d) + " O: " + strconv.Itoa(o) + " S: " + strconv.Itoa(s) + "   OCHO: " + strconv.Itoa(ocho))
				}
			}
		}
	}
}
