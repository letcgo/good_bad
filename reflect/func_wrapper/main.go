package main

import (
	"fmt"
	"reflect"
	"runtime"
	"time"
)

func funcWrap(f interface{}) interface{} {
	rf := reflect.TypeOf(f)
	if rf.Kind() != reflect.Func {
		panic("expects a function")
	}
	vf := reflect.ValueOf(f)
	//[]reflect.Value is runtime args, not called func args
	wrapperF := reflect.MakeFunc(rf, func(in []reflect.Value) []reflect.Value {
		start := time.Now()
		out := vf.Call(in)
		end := time.Now()
		fmt.Printf("calling %s took %v\n", runtime.FuncForPC(vf.Pointer()).Name(), end.Sub(start))
		return out
	})
	return wrapperF.Interface()
}

func timeMe() {
	fmt.Println("starting")
	time.Sleep(1 * time.Second)
	fmt.Println("ending")
}

func timeMeToo(a int) int {
	fmt.Println("starting")
	time.Sleep(time.Duration(a) * time.Second)
	result := a * 2
	fmt.Println("ending")
	return result
}
//demo https://play.golang.org/p/j5kXdqkh-3G
func main() {
	timed := funcWrap(timeMe).(func())
	timed()
	wrongFunc,ok := funcWrap(timeMeToo).(func(int64) int64)
	fmt.Println(wrongFunc,  ok)//wrongFunc is nil,  wrongFunc is a func value, not called
	timedToo := funcWrap(timeMeToo).(func(int) int)
	fmt.Println(timedToo(2))
}
