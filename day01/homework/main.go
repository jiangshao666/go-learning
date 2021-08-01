package main

import "fmt"

func main() {
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			if j == i {
				fmt.Printf("%d X %d = %d\n", i, j, i*j)
			} else {
				fmt.Printf("%d X %d = %d ", i, j, i*j)
			}
		}
	}
}
