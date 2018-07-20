package main

import (
	"fmt"

	"runtime"
	"time"
)
func say(s string) {

	for i := 0; i < 2; i++ {
		fmt.Println(s)
		runtime.Gosched()

	}
}

func main() {
	go say("world")
	go say("hello")
	time.Sleep(1*time.Second)
}
