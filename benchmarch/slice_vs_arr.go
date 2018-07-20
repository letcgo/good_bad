package main

import (
	"time"
)

func modify(sls []int) {
    sls[0] = 90
}

func main() {  
    a := []int{89, 90, 91, 234,23,23,23,234,23,423,423,432,423,4,23,423,423,423}
	benchmark(a[:])
	benchmark(a)
}

func benchmark(a []int) {
	start := time.Now()
	for i := 0; i < 999999999; i++ {
		modify(a)
	}
	println("elapsed nanoseconds:  ", time.Now().Sub(start))
}
