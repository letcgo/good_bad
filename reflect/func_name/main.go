package main

import "fmt"
import "reflect"
import "runtime"

// demo https://play.golang.org/p/0wUXD0HqDab
func main() {
	name := runtime.FuncForPC(reflect.ValueOf(main).Pointer()).Name()
	fmt.Println("Name of function : " + name)
}