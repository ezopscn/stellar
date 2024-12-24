package utils

// 判断字符串是否在切片中
func IsStringInSlice(str string, slice []string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}

func IsUintInSlice(num uint, slice []uint) bool {
	for _, s := range slice {
		if s == num {
			return true
		}
	}
	return false
}
