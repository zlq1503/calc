package calc

// Add 返回两个整数和的函数
func Add(x, y int) int {
	return x + y
}

// Add 返回两个整数差的函数
func Subtract(x, y int) int {
	return x - y
}

// Add 返回两个整数相乘的函数
func Multiply(x, y int) int {
	return x * y
}

// Add 返回两个整数相除的函数
func Divide(x, y int) float32 {

	z := float32(x) / float32(y)

	return z
}
