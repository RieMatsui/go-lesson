package main

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
)


func TestFileOpen(t *testing.T) {
	filename := "test.txt"

	contents := OpenFile(filename)

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

		contents := OpenFile("no_file.txt")
		if contents != "" {
			t.Fatalf("結果が期待通りではありませんでした: %v != true", contents)
		}
	}()
}

func TestFileOpenNoText(t *testing.T) {
	filename := "test3.txt"

	contents := OpenFile(filename)

	if contents != "" {
		t.Fatalf(`Contents is not "" but it is %s`, contents)
	}
}

func TestFileWite(t *testing.T) {

	ok := WiteFile("test5.txt")

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

		ok := WiteFile("")
		if ok {
			t.Fatalf("結果が期待通りではありませんでした: %v != true", ok)
		}
	}()
}

func TestOpenfileByDetailSetting(t *testing.T) {
	ok := OpenfileByDetailSetting("notes.txt")
	if!ok {
		t.Fatalf("結果が期待通りではありませんでした: %v != true", ok)
	}
}

func TestOpenfileByDetailSettingNG(t *testing.T) {
	ok := OpenfileByDetailSetting("notesNoFile.txt")
	if ok {
		t.Fatalf("結果が期待通りではありませんでした: %v != true", ok)
	}
}

func ExampleMain() {
	main()
	// Output:
	// TEST
	// true
	// true
}
