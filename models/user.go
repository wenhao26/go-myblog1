package models

import (
	"gorm.io/gorm"

	"blog/utils"
	"blog/utils/errmsg"
)

type User struct {
	gorm.Model
	Username string
	Password string
	Role     int
}

var user User

// 检查用户是否存在
func CheckUser(name string) bool {
	db.Select("id").Where("username=?", name).First(&user)

	return user.ID > 0
}

// 创建用户
func CreateUser(user *User) int {
	err := db.Create(user).Error
	if err != nil {
		return 0
	}
	return int(user.ID)
}

// 校验登录
func CheckLogin(username, password string) (User, int) {
	db.Where("username=?", username).First(&user)

	if user.ID == 0 {
		return user, errmsg.UserNotExist
	}
	if !utils.CheckPassword(user.Password, password) {
		return user, errmsg.UserLoginWrong
	}
	return user, 0
}
