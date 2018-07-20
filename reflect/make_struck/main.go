package main

import (
	"fmt"
	"reflect"
)

func MakeStruct(vals ...interface{}) interface{} {
	var sfs []reflect.StructField
	for k, v := range vals {
		t := reflect.TypeOf(v)
		sf := reflect.StructField{
			Name: fmt.Sprintf("F%d", (k + 1)),
			Type: t,
		}
		sfs = append(sfs, sf)
	}
	st := reflect.StructOf(sfs)
	so := reflect.New(st)
	return so.Interface()
}
//demo https://play.golang.org/p/lJiTP6vYYN
func main() {
	s := MakeStruct(0, "", []int{})
	// this returned a pointer to a struct with 3 fields:
	// an int, a string, and a slice of ints
	// but you can't actually use any of these fields
	// directly in the code; you have to reflect them
	sr := reflect.ValueOf(s)

	// getting and setting the int field
	fmt.Println(sr.Elem().Field(0).Interface())
	sr.Elem().Field(0).SetInt(20)
	fmt.Println(sr.Elem().Field(0).Interface())

	// getting and setting the string field
	fmt.Println(sr.Elem().Field(1).Interface())
	sr.Elem().Field(1).SetString("reflect me")
	fmt.Println(sr.Elem().Field(1).Interface())

	// getting and setting the []int field
	fmt.Println(sr.Elem().Field(2).Interface())
	v := []int{1, 2, 3}
	rv := reflect.ValueOf(v)
	sr.Elem().Field(2).Set(rv)
	fmt.Println(sr.Elem().Field(2).Interface())
}
