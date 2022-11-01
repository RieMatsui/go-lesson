package main

import (
	"fmt"
	"os"
)

func main() {
	// Args
	fmt.Println(os.Args[0])
	fmt.Println(os.Args[1])
	fmt.Println(os.Args[2])
	fmt.Println(os.Args[3])

	// os.Argsの要素数を表示
	fmt.Printf("length=%d\n", len(os.Args))

	for i, v := range os.Args {
		fmt.Println(i, v)
	}
}

