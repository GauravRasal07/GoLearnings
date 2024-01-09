package main

import (
	"fmt"
)

func main(){
	nums := []int{0,1,2,3,4,5,6,7,8,9}

	for _, num := range nums {
		fmt.Println(isOddOrEven(num))
	}
}

func isOddOrEven(num int) string {
	if num % 2 == 0 {
		return "EVEN"
	} else {
		return "ODD"
	}
}