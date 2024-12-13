package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://github.com/SudhirYadav5678"

func main() {
	fmt.Println("Web")
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	//fmt.Printf("tyle is %T", res)
	//fmt.Println(res)
	defer res.Body.Close()

	datbytes, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(datbytes))
}
