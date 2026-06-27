package main

import "fmt"

func main() {
	numbers := []int{10, 20, 30, 40, 50}
	fmt.Println(numbers)
	numbers = append(numbers, 60)
	fmt.Println(numbers)
	fmt.Println(numbers[1:3])
	fmt.Println(len(numbers))
	fmt.Println(cap(numbers))
}
