package main

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestFileOpen(t *testing.T) {
	filename := "test.txt"

	contents := FileOpen(filename)

	if contents != "TEST" {
		t.Fatalf(`Contents is not "Hello World!" but it is %s`, contents)
	}
}

func TestFileOpenNo(t *testing.T) {
	setUp := func() {
		log.SetOutput(ioutil.Discard) // /dev/nullに出力
	}
	tearDown := func() {
		println("tearDown")
		log.SetOutput(os.Stdout) // 出力先を戻す
	}

	func() {
		setUp()
		defer tearDown()

		contents := FileOpen("no_file.txt")
		if contents != "" {
			t.Fatalf("結果が期待通りではありませんでした: %v != true", contents)
		}
	}()
}

func TestFileOpenNoText(t *testing.T) {
	filename := "test3.txt"

	contents := FileOpen(filename)

	if contents != "" {
		t.Fatalf(`Contents is not "" but it is %s`, contents)
	}
}

func TestFileWite(t *testing.T) {

	ok := FileWite("test5.txt")

	if !ok {
		t.Fatalf("結果が期待通りではありませんでした: %v != true", ok)
	}
}

func TestFileWiteNG(t *testing.T) {
	setUp := func() {
		log.SetOutput(ioutil.Discard) // /dev/nullに出力
	}
	tearDown := func() {
		println("tearDown")
		log.SetOutput(os.Stdout) // 出力先を戻す
	}

	func() {
		setUp()
		defer tearDown()

		ok := FileWite("")
		if ok {
			t.Fatalf("結果が期待通りではありませんでした: %v != true", ok)
		}
	}()
}

