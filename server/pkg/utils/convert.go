package utils

import "strconv"

// 将字符串转换为 uint
func StringToUint(str string) uint {
	val, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return 0
	}
	return uint(val)
}
