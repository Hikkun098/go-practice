package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Println("1に1を足すこれを10回行う:", i)
	}

	count := 0
	for count < 20 {
		fmt.Println(count)
		count++
	}

	fruits := []string{"apple", "orange", "banana"}
	for index, value := range fruits {
		fmt.Printf("index: %d, value: %s\n", index, value)
	}
}
