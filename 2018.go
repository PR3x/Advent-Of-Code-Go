package main

import "fmt"

func main() {
	day1 := make(chan int)
	go Day1(day1)
	fmt.Println("d1p1:", <-day1)
	fmt.Println("d1p2:", <-day1)
}
