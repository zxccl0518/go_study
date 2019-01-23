package main

import "fmt"

func main() {
	// continue 的演示。

	// lable2:
	for i := 1; i < 10; i++ {
		for j := 0; j < 5; j++ {
			if j == 2 {
				continue
			}
			fmt.Println("j = ", j)
		}
	}
}
