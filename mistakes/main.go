package main

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
)

const debug = false

type MyStruct struct {

}

func main() {
	var buf *bytes.Buffer //原来是var buf *bytes.Buffer
	f(buf)
	buf = new(bytes.Buffer)
	f(buf)
	var scalar float64
	check("scalar", scalar)
	//if scalar == nil{ //cannot convert nil to type float64
	//}
	//if scalar == (float64)(nil) {//cannot convert nil to type float64
	//	fmt.Println("surprise!!! scalar type is float64 and value is nil!")
	//}
	if scalar == 0 {
		fmt.Println("surprise!!! scalar type is float64 and value is 0 !")
	}
	var ptrStruct *MyStruct
	check("ptrStruct", ptrStruct)
	if ptrStruct == (*MyStruct)(nil) {
		fmt.Println("surprise!!! ptrStruct type is *MyStruct and value is nil!")
	}

	var varStruct MyStruct
	check("varStruct", varStruct)
	//nil is useful for pointer, not struct nor scalar
	//if varStruct == (MyStruct)(nil) {//cannot convert nil to type MyStruct
	//	fmt.Println("surprise!!! varStruct type is *MyStruct and value is nil!")
	//}
	//if varStruct == nil {//cannot convert nil to type MyStruct
	//	fmt.Println("surprise!!! varStruct type is *MyStruct and value is nil!")
	//}
}
func f(buf io.Writer) {
	if buf == (*bytes.Buffer)(nil) {
		fmt.Println("surprise!!! type is *bytes.Buffer and value is nil!")
	}
	if buf != nil {
		fmt.Println("surprise!!! value is not nil")
	}
	check("buf", buf)
}

func check(name string , v interface{}){
	fmt.Printf("%s is %#v, type is %s, v is nil %b  \n", name, v, reflect.TypeOf(v).String(), v==nil)
}