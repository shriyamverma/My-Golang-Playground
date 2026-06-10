package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Circle struct {
	radius float64
}

type Rectangle struct {
	length, width float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r Rectangle) Area() float64 {
	return r.length * r.width
}

func calculateArea(shape Shape) float64 {
	return shape.Area()
}

func main() {
	circle := Circle{radius: 5}
	rectangle := Rectangle{length: 5, width: 5}

	fmt.Println("Circle area: ", calculateArea(circle))
	fmt.Println("Rectangle area: ", calculateArea(rectangle))

	fmt.Println("-------------------------")

	mysteryBox := interface{}(10)
	describeValue(mysteryBox)

	retrievedInt, ok := mysteryBox.(int)
	if ok {
		fmt.Println("Retrieved value is: ", retrievedInt)
	} else {
		fmt.Println("Retrieved value is not an int")
	}

	fmt.Println("-------------------------")

}

func describeValue(value interface{}) {
	fmt.Printf("Type: %T, Value: %v\n", value, value)
}
