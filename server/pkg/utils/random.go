package utils

import "math/rand"

// 生成指定长度的随机字符串
func RandString(n int, letters string) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = rune(letters[rand.Intn(len(letters))])
	}
	return string(b)
}

// 生成指定长度的数字字符串
func RandNumberString(n int) string {
	return RandString(n, "1234567890")
}
