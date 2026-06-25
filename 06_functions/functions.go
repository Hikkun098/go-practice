package main

import "fmt"

func add(a int, b int) int {
    return a + b
}

func divide(a int, b int) (int, int) {
    quotient := a / b
    remainder := a % b
    return quotient, remainder
}

func main() {
    quotient, remainder := divide(10, 3)
    fmt.Println(quotient, remainder)
    greet("Alice")
}

func greet(name string) {
    fmt.Printf("Hello, %s!\n", name)
}