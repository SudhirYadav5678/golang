package main

import "fmt"

// no var style only allowed in any function
//hello:=8045

//Captical letter consider as the public function
//const Letter string= "Hello"

func main() {
	var username string = "Sudhir"
	fmt.Println(username)
	fmt.Printf("This is your %T \n", username)

	//no var style
	// no var style only allowed in any function
	number := 499
	fmt.Println(number)
}
