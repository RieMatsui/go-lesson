package main

import (
	"fmt"
	"time"
)

func main() {

	t2 := time.Date(2020, 6, 20, 10, 16, 19, 1, time.Local)
	fmt.Println(t2)
	fmt.Println(t2.Year())
	fmt.Println(t2.YearDay())
	fmt.Println(t2.Month())
	fmt.Println(t2.Weekday())
	fmt.Println(t2.Day())
	fmt.Println(t2.Hour())
	fmt.Printf("%T\n", time.Hour)
	fmt.Println(t2.Minute())
	fmt.Println(t2.Second())
	fmt.Println(t2.Nanosecond())
	fmt.Println(t2.Zone())
	printTimeDuration()
	changeTime()
	addTime(t2, 2, 15)
}

func changeTime() {
	d, _ := time.ParseDuration("2h30m")
	fmt.Println(d)
}

func addTime(t time.Time,minute time.Duration, second time.Duration) {
	t.Add(minute*time.Minute + second*time.Second)
	fmt.Println(t)
}

func printTimeDuration() {
	fmt.Println(time.Hour)
	fmt.Println(time.Minute)
    fmt.Println(time.Second)
	fmt.Println(time.Millisecond)
	fmt.Println(time.Microsecond)
	fmt.Println(time.Nanosecond)
}


func conpairTime (){
	t1 := time.Date(2020, 6, 20, 10, 16, 19, 1, time.Local)
    t2 := time.Date(2022, 11, 1, 11, 22, 20, 3, time.Local)

	d2 := t2.Sub(t1)
	fmt.Println(d2)

	fmt.Println(t2.Before(t1))
	fmt.Println(t2.After(t1))

	fmt.Println(t1.Before(t2))
	fmt.Println(t1.After(t2))
}

func sleep(){
	time.Sleep(5 * time.Second)
	fmt.Println("5秒停止後表示")
}