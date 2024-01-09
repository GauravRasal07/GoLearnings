package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("https://randomuser.me/api/?results=1`")

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	fmt.Println("The response code is:\n", resp.StatusCode)

	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println("The response body is:\n", string(bs))		//First way of reading body

	io.Copy(os.Stdout, resp.Body)								//Second way of reading body
}