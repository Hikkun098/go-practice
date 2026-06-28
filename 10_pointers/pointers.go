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
	fmt.Println("値:", num)
	fmt.Println("アドレス:", &num)
  double(num)
	fmt.Println("doubleの後:", num)
	doublePointer(&num)
	fmt.Println("doublePointerの後:", num)
	fmt.Println("アドレス:", &num)
	fmt.Println("関数の外:", num)
}
