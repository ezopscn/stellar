package utils

import (
	"regexp"
	"strconv"
	"strings"
)

// 验证邮箱
func IsEmail(email string) bool {
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	return regexp.MustCompile(regex).MatchString(email)
}

// 验证手机号
func IsPhoneNumber(number string) bool {
	regex := `1[3-9]\d{9}$`
	return regexp.MustCompile(regex).MatchString(number)
}

// 验证 QQ 号
func IsQQNumber(number string) bool {
	regex := `[1-9]([0-9]){5,11}`
	return regexp.MustCompile(regex).MatchString(number)
}

// 验证身份证
func IsIDCard(id string) bool {
	regex1 := `^[1-9]\d{5}\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\d{3}$`
	regex2 := `^[1-9]\d{5}(18|19|([23]\d))\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\d{3}[0-9Xx]$`
	return regexp.MustCompile(regex1).MatchString(id) || regexp.MustCompile(regex2).MatchString(id)
}

// 验证 IPv4 要求
func isValidIPAddressRange(ip string) bool {
	parts := strings.Split(ip, ".")
	for _, part := range parts {
		num, _ := strconv.Atoi(part)
		if num < 0 || num > 255 {
			return false
		}
	}
	return true
}

// 验证 IPv4
func IsIPv4(ip string) bool {
	regex := `^(?:[0-9]{1,3}\.){3}[0-9]{1,3}$`
	return regexp.MustCompile(regex).MatchString(ip) && isValidIPAddressRange(ip)
}

// 验证 IPv6
func IsIPv6(ip string) bool {
	regex := `^((?:[0-9a-fA-F]{1,4}:){7}[0-9a-fA-F]{1,4}|::|((?:[0-9a-fA-F]{1,4}:){1,7}:)|((?:[0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4})|((?:[0-9a-fA-F]{1,4}:){1,5}((?::[0-9a-fA-F]{1,4}){1,2}))|((?:[0-9a-fA-F]{1,4}:){1,4}((?::[0-9a-fA-F]{1,4}){1,3}))|((?:[0-9a-fA-F]{1,4}:){1,3}((?::[0-9a-fA-F]{1,4}){1,4}))|((?:[0-9a-fA-F]{1,4}:){1,2}((?::[0-9a-fA-F]{1,4}){1,5}))|[0-9a-fA-F]{1,4}:((?::[0-9a-fA-F]{1,4}){1,6})|:((?::[0-9a-fA-F]{1,4}){1,7}|:))$`
	return regexp.MustCompile(regex).MatchString(ip)
}

// 验证端口号
func IsPort(port string) bool {
	regex := `^([0-9]{1,5})$`
	if !regexp.MustCompile(regex).MatchString(port) {
		return false
	}
	num, _ := strconv.Atoi(port)
	return num >= 1 && num <= 65535
}

// 验证 Mac 地址
func IsMacAddress(mac string) bool {
	regex := `^([0-9a-fA-F]{2}[:-]){5}([0-9a-fA-F]{2})$`
	return regexp.MustCompile(regex).MatchString(mac)
}
