package main

import "fmt"

func main() {

	fmt.Println("This is the entry point")

	for i := 1; i <= 50; i++ {
		if i%2 == 0 {
			fmt.Println("", i)

		} else {
			fmt.Print(i)
		}

	}

	//else
	//fmt.Print(i)
	fmt.Println()
	fmt.Println("This is the exit point")
}
