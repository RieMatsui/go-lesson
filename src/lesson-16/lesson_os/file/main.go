package main

import (
	"fmt"
	"log"
	"os"
)

func OpenFile(file string) (contents string) {
	f, err := os.Open(file)
	if err != nil {
		log.Println("ファイルがありません:", err)
		return ""
	}
	defer f.Close()

	buf := make([]byte, 1024)
	for {
		n, _ := f.Read(buf)

		if n == 0 {
			break
		}
		return string(buf)[:n]
	}
	return ""
}

func WiteFile(file string) (ok bool) {
	f, err := os.Create(file)
	if err != nil {
		log.Println(err)
		return false
	}
	f.Write([]byte("Hello\n"))

	f.WriteAt([]byte("Golang"), 6)

	f.Seek(0, os.SEEK_END)

	f.WriteString("Yeah")
	return true
}

func OpenfileByDetailSetting(file string) bool {
	f, err := os.OpenFile(file, os.O_RDONLY, 0755)
	if err != nil {
		log.Println(err)
		return false
	}
	if err := f.Close(); err != nil {
		log.Println(err)
		return false
	}
	return true
}

func main() {
	contents := OpenFile("test.txt")
	fmt.Println(contents)

	fg_write := WiteFile("write_test.txt")
	fmt.Println(fg_write)

	fg_open := OpenfileByDetailSetting("notes.txt")
	fmt.Println(fg_open)
}
