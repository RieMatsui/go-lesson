package main

import "fmt"

func main() {

	defer func ()  {
		if x := recover(); x != nil {
			fmt.Printf("Error: %v\n", x)
		}
	}()

	panic("runtime error")
	// * 下記は実行されない
	// * fmt.Println("start")
}
