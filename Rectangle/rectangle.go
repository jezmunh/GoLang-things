package main

import "fmt"

type rectangle struct {
	width, height int
}

func (rect *rectangle) area() int {
	return rect.width * rect.height
}

func (rect *rectangle) perimeter() int {
	return 2 * (rect.width + rect.height)
}

func main() {
	rect1 := rectangle{width: 4, height: 2}

	fmt.Println("Area of the rectangle:", rect1.area())
	fmt.Println("Perimeter of the rectangle:", rect1.perimeter())
}
