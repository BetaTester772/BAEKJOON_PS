package main

import (
	"fmt"
)

func main() {
	var a, b int
	for {
		_, err := fmt.Scanf("%d %d", &a, &b)
		if err != nil {
			break
		}
		fmt.Printf("%d\n", a+b)
	}
}
