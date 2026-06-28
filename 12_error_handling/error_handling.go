package main

import (
    "errors"
    "fmt"
)

func divide(a, b float64) (float64, error) {
    // bが0ならエラーを返す
    // ヒント: errors.New("エラーメッセージ")
    // 正常なら計算結果とnilを返す
		if b == 0 {
			return 0, errors.New("0で割ることはできません。")
		}
		return a / b, nil
}

func main() {
	firstResult, err := divide(10, 3)
	if err != nil {
		fmt.Println("エラー:", err)
	} else {
		fmt.Println("結果:", firstResult)
	}

	secondResult, err2 := divide(10, 0)
	if err2 != nil {
		fmt.Println("エラー:", err2)
	} else {
		fmt.Println("結果:", secondResult)
	}
}