package main

import "fmt"

type ICar interface{
	drive()
}
type BenzA200L struct{
}
func (t BenzA200L)drive(){
	fmt.Println(t, "driving BenzA200L")
}

type BenzA200 struct{
}
func (t *BenzA200)drive() {
	fmt.Println(t, "driving BenzA200")
}
func main()  {
	var a200l ICar  = BenzA200L{}
	a200l.drive()
	var a200 ICar = new(BenzA200)
	a200.drive()

}
