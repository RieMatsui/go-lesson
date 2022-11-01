package main

import (
	"log"
	"os"
)

func FileOpen(file string) (contents string){
	f, err := os.Open(file)
	if err != nil {
		log.Println("ファイルがありません:" , err)
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

func FileWite(file string)(ok bool){
	f, err := os.Create(file)
    if err!= nil {
        log.Println(err)
		return false
    }
	f.Write([]byte("Hello\n"))

	f.WriteAt([]byte("Golang"), 6)

	f.Seek(0, os.SEEK_END)

	f.WriteString("Yeah")
	return true
}

// func main() {
// 	FileOpen("test.txt")
// 	FileWite("test2.txt")
// }
