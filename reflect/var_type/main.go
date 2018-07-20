package main

import (
	"fmt"
	"reflect"
)

type Ab struct {
}

func getType(myvar interface{}) string {
	return reflect.TypeOf(myvar).Name()
}
//demo https://play.golang.org/p/zvAmNp1ALr7
func main() {
	fmt.Println("Hello, playground")

	tst := "string"
	tst2 := 10
	tst3 := 1.2
	tst4 := new(Ab)

	fmt.Println(getType(tst))
	fmt.Println(getType(tst2))
	fmt.Println(getType(tst3))
	fmt.Println(getType(tst4))

}
