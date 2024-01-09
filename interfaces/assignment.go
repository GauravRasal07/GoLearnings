package main

import "fmt"

type shape interface{
	getArea() float64
}

type triangle struct{
	base float64
	height float64
}

type square struct{
	side float64
}

func (t triangle) getArea() float64 {
	area := 0.5 * t.base * t.height
	return area
}

func (s square) getArea() float64 {
	area := s.side * s.side
	return area
}

func main(){
	tri := triangle{
		base: 5,
		height: 6,
	}

	areaOfTriangle := tri.getArea()

	squ := square{
		side: 5,
	}

	areaOfSquare := squ.getArea()

	fmt.Println("The area of triangle is: ", areaOfTriangle)
	fmt.Println("The area of square is: ", areaOfSquare)
}