package main

import "fmt"

func Add1[T int | float64](a T, b T) T {
	return a + b
}

type Num interface {
	int | int32 | int64 | float32 | float64
}

func Add2[T Num](a T, b T) T {
	return a + b
}

type UserID int

// Add3 ~(tilde) allows you to use any type that is alias for this int
func Add3[T ~int | float64](a T, b T) T {
	return a + b
}

func main() {
	result1 := Add1(1.1, 2.2)
	fmt.Println(result1)

	result2 := Add2(1, 2.2)
	fmt.Println(result2)

	a := UserID(1)
	b := UserID(2)
	result3 := Add3(a, b)
	fmt.Println(result3)
}
