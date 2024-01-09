package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	// "strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}

	fmt.Println("The server is listening!")

	defer li.Close()

	for {
		conn, err := li.Accept()

		if err != nil {
			log.Panic(err)
		}

		go handle(conn)
	}

}

func handle(conn net.Conn) {
	// text := "This is to test something which is failing because of something!!! :)"

	// scanner := bufio.NewScanner(strings.NewReader(text))
	scanner := bufio.NewScanner(conn)
	
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
	}

	defer conn.Close()


	fmt.Println("*****Your code ever got here*****")
}
