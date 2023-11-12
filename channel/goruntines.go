package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 100; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	f("Hi")
	f("mass")
	go f("Natty Hi")
	go f("massage2")
	time.Sleep(5 * time.Second)
}
