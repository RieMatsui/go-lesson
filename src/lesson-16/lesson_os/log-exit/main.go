package main

import (
	"fmt"
	"log"
	"os"
)

var osExit = os.Exit

func Die() {
    osExit(1)
}

func DeferFunc() {
	defer func() {
		fmt.Println("defer")
	}()
	osExit(1)
}

func LogFunc() {
	_, err := os.Open("A.text")
	if err != nil {
		log.Fatalln(err)
	}
}