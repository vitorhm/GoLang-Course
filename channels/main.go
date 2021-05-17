package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://udemy.com",
	}

	for _, link := range links {
		checkLink(link)
	}
}

func checkLink(l string) {
	_, error := http.Get(l)

	if error != nil {
		fmt.Println(error)
	}
}
