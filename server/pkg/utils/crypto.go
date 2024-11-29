package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// 密码加密
func CryptoPassword(password string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash)
}

// 密码验证
func ComparePassword(hash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// 手机号加密
func HidePhoneNumber(phone string) string {
	return fmt.Sprintf("%s****%s", phone[:3], phone[7:])
}
