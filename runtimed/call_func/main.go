package main

import "fmt"

func MyFunc1() {
	fmt.Println("I am Masood")
}

func MyFunc2() {
	fmt.Println("I am a programmer")
}

func MyFunc3() {
	fmt.Println("I want to buy a car")
}
//demo https://play.golang.org/p/WNBg4bezDBu
func main() {
	registry := map[string]func(){
		"MyFunc1": MyFunc1,
		"MyFunc2": MyFunc2,
		"MyFunc3": MyFunc3,
	}
	for k := 1; k <= 3; k++ {
		registry[fmt.Sprintf("MyFunc%d", k)]()
	}
}
