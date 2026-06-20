package main

import "fmt"

func main() {
	// ===== var宣言（型を明示する書き方） =====
	var name string = "Hidehiro"
	var age int = 30
	var height float64 = 171.0
	var isEngineer bool = true

	fmt.Println("名前:", name)
	fmt.Println("年齢:", age)
	fmt.Println("身長:", height)
	fmt.Println("エンジニア?:", isEngineer)

	// ===== 短縮宣言（:= を使う書き方） =====
	// 型を書かなくても、右辺の値から自動で推論してくれる
	company := "SIer企業"
	experience := 1 // int になる
	coffee := 3.5   // float64 になる

	fmt.Println("会社:", company)
	fmt.Println("経験年数:", experience)
	fmt.Println("今日のコーヒー杯数:", coffee)

	// ===== var と := の使い分け =====
	// var は「初期値なしで宣言したい」ときに便利
	// 初期値を入れないと、型ごとのゼロ値が自動で入る
	var score int      // 0 が入る
	var message string // "" (空文字) が入る
	var active bool    // false が入る

	fmt.Println("スコア:", score)
	fmt.Println("メッセージ:", message)
	fmt.Println("アクティブ?:", active)

	// ===== 変数の値を更新する =====
	// 一度宣言した変数は = で更新（:= は使わない）
	score = 100
	message = "Go楽しい！"
	active = true

	fmt.Println("--- 更新後 ---")
	fmt.Println("スコア:", score)
	fmt.Println("メッセージ:", message)
	fmt.Println("アクティブ?:", active)

	// ===== まとめて宣言 =====
	var (
		x int    = 10
		y int    = 20
		z string = "合計"
	)
	fmt.Println(z, ":", x+y)
}