package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Hello web")
	//PerformGet()
	PerformPost()
}

func PerformGet() {
	const myurl = "https://github.com/SudhirYadav5678?tab=repositories"

	res, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	fmt.Println("status", res.Status)
	fmt.Println("length", res.ContentLength)
	//Hello web status 200 OK length -1

	content, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content))
}

func PerformPost() {
	const myurl = "http://localhost:8000"
	reqBody := strings.NewReader(`
          {"name":"sudhir",
            "hello":"hello"}
    `)
	res, err := http.Post(myurl, "application/json", reqBody)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	content, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content))
}

func PerformPostForm() {
	const myurl = "http://localhost:8000"
	data := url.Values{}
	data.Add("firstName", "hitesh")

	res, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	content, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content))
}
