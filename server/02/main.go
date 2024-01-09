package main 

import (
	"fmt"
	"net"
	"bufio"
	"log"
	"time"
)

func main(){
	li, err := net.Listen("tcp", ":8080")
	
	if err != nil{
		log.Panic(err)
	}

	defer li.Close()

	for {
		conn, err := li.Accept()

		if err != nil{
			fmt.Println("You got an error: ", err)
			log.Panic(err)
		}

		go handle(conn)
	}
}

func handle(con net.Conn){
	err := con.SetDeadline(time.Now().Add(10 * time.Second))
	if err != nil{
		fmt.Println("The connection timed out!!!")
	}

	scanner := bufio.NewScanner(con)

	for scanner.Scan(){ 
		ln := scanner.Text()
		fmt.Println("You got here: ", ln)
		fmt.Fprintf(con, "You just said: %s\n", ln)
	}

	defer con.Close()

	fmt.Println("***** Your code got here *****")
}