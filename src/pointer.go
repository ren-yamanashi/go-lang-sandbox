package main

func _pointer() {
	var num1 int
	var point1 *int

	point1 = &num1 // point1 に num1 のポインタを設定 (`&` でポインタを参照)
	*point1 = 10   // ポインタの中身に 10 を代入 (`*` でポインタの中身を参照して値を代入)

	println("num1:", num1)       // num1: 10
	println("point1:", *point1) // point1: 10

	num2 := 0
	passReference(&num2) // num2 のポインタを渡す
	println("num2:", num2) // num2: 20
}

// 参照を受け取って値を変更
func passReference(p *int) {
	*p = 20 // ポインタの中身に 20 を代入
}