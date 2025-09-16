package user

import (
	"AILN/app/common"
)

type User struct {
	ID          uint   `json:"id" gorm:"id"`
	Username    string `json:"username" gorm:"username"`
	Password    string `json:"password" gorm:"password"`
	Description string `json:"description" gorm:"description"`
}

func Create(u *User) error {
	return common.DB.Create(u).Error
}

func ExistUP(username string, password string) bool {
	return common.DB.Table("user").Where("username = ? AND password = ?", username, password).First(&User{}).Error == nil
}

func FindOneByUP(username string, password string) (user *User, err error) {
	user = new(User)
	err = common.DB.Table("user").Where("username = ? AND password = ?", username, password).First(user).Error
	return
}

func FindInPage(page uint, pageSize uint) (users []*User, err error) {
	err = common.DB.Table("user").Offset(int((page - 1) * pageSize)).Limit(int(pageSize)).Find(&users).Error
	return users, err
}
