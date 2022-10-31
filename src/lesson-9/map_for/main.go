package main

import "fmt"

func main() {
	m := map[string]int {
		"Apple" : 100,
		"Banana" : 200,
	}

	for k, v:= range m {
		fmt.Println(k, v)
	}

	for _, v:= range m {
		fmt.Println(v)
	}

	for i := range m {
		fmt.Println(i)
	}
}
