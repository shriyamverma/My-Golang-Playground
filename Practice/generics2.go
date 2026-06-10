package main

import "fmt"

// input: [1, 2, 3], (n) => n * 2
// output: [2, 4, 6]

func MapValues(values []int, mapFunc func(int) int) []int {
	var newValues []int
	for _, value := range values {
		newValue := mapFunc(value)
		newValues = append(newValues, newValue)
	}
	return newValues
}

func main() {
	result := MapValues([]int{1, 2, 3}, func(n int) int {
		return n * 2
	})
	fmt.Printf("result: %v\n", result)
}
