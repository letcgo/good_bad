package main

import (
	"fmt"
	"reflect"
)
//demo https://play.golang.org/p/ZMFF3a1sO3h
func main() {
	// declaring these vars, so I can make a reflect.Type
	intSlice := make([]int, 0)
	mapStringInt := make(map[string]int)

	// here are the reflect.Types
	sliceType := reflect.TypeOf(intSlice)
	mapType := reflect.TypeOf(mapStringInt)

	// and here are the new values that we are making
	intSliceReflect := reflect.MakeSlice(sliceType, 0, 0)
	mapReflect := reflect.MakeMap(mapType)

	// and here we are using them
	v := 10
	rv := reflect.ValueOf(v)
	intSliceReflect = reflect.Append(intSliceReflect, rv)
	intSlice2 := intSliceReflect.Interface().([]int)
	fmt.Println(intSlice2)

	k := "hello"
	rk := reflect.ValueOf(k)
	mapReflect.SetMapIndex(rk, rv)
	mapStringInt2 := mapReflect.Interface().(map[string]int)
	fmt.Println(mapStringInt2)
}
