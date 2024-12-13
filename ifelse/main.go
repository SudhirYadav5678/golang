package main

import "fmt"

func main() {
	result := 34
	if result < 20 {
		fmt.Println(result)
	} else {
		fmt.Println("Greater")
	}

	defer fmt.Println("Hello") //last
	defer fmt.Println("Yadav") //sec
	fmt.Println("Sudhir")      //first
}
