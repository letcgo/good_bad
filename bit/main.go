package main

import (
	_ "sdk"
	"strconv"
	"fmt"
)

func main() {
	str  := "abc"
	for i:=0; i<100; i++{
		a := i<<2
		//b := 2<<(int)i
		go func(i, a int) {
			//println(str+strconv.Itoa(i)+" "+strconv.Itoa(a))
			fmt.Println(str+strconv.Itoa(i)+" "+strconv.Itoa(a))
		}(i,a)
	}
}
