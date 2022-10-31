package main

import (
	"fmt"

	"go-lesson/src/lesson-13/foo"
)

func appName() string {
	const AppName = "GoApp"
	var Version string = "1.6"
	return AppName + " " + Version
}

func Do(s string)(b string) {
	b = s
	{
		b := "BBBB"
		fmt.Println(b)
	}
	fmt.Println(b)
	return b
}

func main() {
	fmt.Println(foo.Max)
	fmt.Println(foo.RetuernMin())

	fmt.Println(appName())

	fmt.Println(Do("AAA"))
}
