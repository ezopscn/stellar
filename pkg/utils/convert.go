package utils

import "strconv"

func StringToUint(str string) uint {
	val, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return 0
	}
	return uint(val)
}
