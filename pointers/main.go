package main

import "fmt"

func main() {
	fmt.Println("Pointers")
	//var one int = 2
	var pointers *int
	fmt.Println(pointers)

	myMuner := 34
	var ptr *int = &myMuner
	fmt.Println(ptr)
	fmt.Println(*ptr)
	*ptr = *ptr + 3
	fmt.Println(*ptr)
	fmt.Println(myMuner)

}
