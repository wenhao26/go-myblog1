package utils

import (
	"golang.org/x/crypto/bcrypt"
)

//生成密码串
func GenPassword(password string) string {
	result, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return ""
	}
	return string(result)
}

// 验证密码串
// password 数据库存储的加密密串
// verifyPassword 用户输入待验证的明文密码
func CheckPassword(password, verifyPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(verifyPassword))
	if err != nil {
		return false
	}
	return true
}
