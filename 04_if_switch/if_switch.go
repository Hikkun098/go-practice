package main

import "fmt"

func main() {
	const firstScore int = 80
  const secondScore int = 90

	if firstScore >= 90 {
		fmt.Println("優秀")
	} else if firstScore >= 60 {
		fmt.Println("合格")
	} else {
		fmt.Println("不合格")
	}

	switch {
  case secondScore >= 90:
    fmt.Println("優秀")
  case secondScore >= 60:
    fmt.Println("合格")
  default:
    fmt.Println("不合格")
  }
}
