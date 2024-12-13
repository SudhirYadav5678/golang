package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("COntents")
	content := "Hellohjashfkjdskfhh"
	file, err := os.Create("./myContent.txt")
	if err != nil {
		panic("File is not create")
	}
	length, err := io.WriteString(file, content)
	if err != nil {
		panic("File is not create")
	}
	fmt.Println("Lentgth", length)
	file.Close()

	readfile("./myContent.txt")
}

func readfile(filename string) {
	//databytes, err := ioutil.ReadFile(filename)
	databytes, err := os.ReadFile(filename)
	if err != nil {
		panic("File is not create")
	}
	fmt.Println(string(databytes))
}
