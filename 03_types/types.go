package main

import "fmt"

func main() {
	var message string = "Thank you!"
	var level int = 100
	var point float64 = 4.5
	var isEngineer bool = true
	total := float64(level) + point

	fmt.Printf("メッセージ: %s\n", message)
	fmt.Printf("レベル: %d\n", level)
	fmt.Printf("ポイント: %f\n", point)
	fmt.Printf("エンジニアですか？: %t\n", isEngineer)
	fmt.Println(total)
	fmt.Printf("float64に変換: %f\n", float64(level))
	fmt.Printf("intに変換: %d\n", int(point))
	result := fmt.Sprintf("レベル: %d", level)
	fmt.Println(result)
}