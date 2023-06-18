package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://instagram.com",
		"http://stackoverflow.com",
		"http://twitter.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkList(link, c)
	}
	
	// give a gap between go routine with time package
	for l := range c {
		go func(link string) {
			time.Sleep(1 * time.Second)
			checkList(link, c)
		}(l)
	}

	// run infinite loop without gap in go routine
	// for {
	// 	go checkList(<-c, c)
	// }
}

func checkList(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "went wrong")
		c <- link
		return
	}
	fmt.Println(link, "good to go")
	c <- link
}