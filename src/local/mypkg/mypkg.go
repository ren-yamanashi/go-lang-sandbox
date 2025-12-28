package mypkg

import "fmt"

func FuncA() {			// 大文字で始まるものは自動的にエクスポートされる
    fmt.Println("Run FuncA!")
}

func funcB() {			// 小文字で始まるものはエクスポートされない
    fmt.Println("Run funcB!")
}