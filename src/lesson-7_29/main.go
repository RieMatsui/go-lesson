package main

import "fmt"

func ReturnFunc() func() {
	return func() {
        fmt.Println("I'am a	function!")
    }
}

func main() {
	f := ReturnFunc()
	f()
}

