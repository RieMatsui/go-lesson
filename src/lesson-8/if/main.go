package main

import "fmt"

func main() {
	a := 0
	if a == 0 {
		fmt.Println("two")
	} else if  a == 1 {
		fmt.Println("one")
	} else {
		fmt.Println("I don't know")
	}

	if b := 100; b == 100 {
		fmt.Println("one handred")
	}

	x := 0
	if x := 2; true {
        fmt.Println(x)
	}
	fmt.Println(x)
}

