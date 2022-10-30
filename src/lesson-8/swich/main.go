package main

import "fmt"



func main() {
	n := 3
	switch n {
	case 1, 2:
		fmt.Println("1 or 2")
	case 3, 4:
		fmt.Println("3 or 4")
	default:
		fmt.Println("I don't know")
	}

	switch n := 2; n {
	case 1, 2:
	fmt.Println("1 or 2")
	case 3, 4:
	fmt.Println("3 or 4")
	default:
	fmt.Println("I don't know")
	}

	n1 := 6
	switch {
	case n1 > 0 && n1 < 4:
		fmt.Println("0 < m < 4")
	case n1 > 3 && n1 < 7 :
		fmt.Println("3 < m < 7")
	}
}
