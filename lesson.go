package main

import "fmt"

func main () {
	for x := 0; x <= 15; x++ {
		if x == 3 {
			continue
		} else if x == 8{
			break
		}
		fmt.Println(x)
	}
}
