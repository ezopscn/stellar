package utils

import "os"

// 文件是否存在
func FileExists(path string) bool {
	info, err := os.Stat(path)
	return !os.IsNotExist(err) && !info.IsDir()
}

// 目录是否存在
func DirExists(path string) bool {
	info, err := os.Stat(path)
	return !os.IsNotExist(err) && info.IsDir()
}
