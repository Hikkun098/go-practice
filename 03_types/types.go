package main

import "fmt"

func main() {
	var message string = "Thank you!"
	var level int = 100
	var point float64 = 4.5
	var isEngineer bool = true

	fmt.Printf("メッセージ:%s", message)
	fmt.Printf("レベル:%d", level)
	fmt.Printf("ポイント:%f", point)
	fmt.Printf("エンジニアですか？:%t", isEngineer)
}