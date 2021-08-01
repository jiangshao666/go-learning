package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				goto endTag
			}
			fmt.Printf("i=%v j=%v \n", i, j)
		}
	}
	return

endTag:
	fmt.Println("over")
}
