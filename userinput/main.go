package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Hello User"
	fmt.Printf(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Reader is reading")

	//comma ok or err Ok
	input, _ := reader.ReadString('\n')
	fmt.Println("Thankd tof", input)
}
