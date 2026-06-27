package main

import "fmt"

func double (n int) {
	n = n * 2
	fmt.Println("関数の中:", n)
}

func doublePointer (n *int) {
	*n = *n * 2
	fmt.Println("関数の中:", *n)
}

func main() {
	num := 10
  double(num)
	doublePointer(&num)
	fmt.Println("関数の外:", num)
}
