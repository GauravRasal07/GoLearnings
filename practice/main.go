package main
 
func main() {
     c := make(chan string)
     c <- "Hi there!"
}