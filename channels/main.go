package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{"https://google.com.in", "https://amazon.in", "https://golang.org", "https://stackoverflow.com",}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	// for {
	// 	go checkLink(<- c, c)
	// }

	for l := range c {
		go func(lnk string){
			time.Sleep(5 * time.Second)
			checkLink(lnk, c)
		}(l)
	}
}

func checkLink(link string, c chan string){
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "seems to be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up and running!")
	c <- link
}
