package user

import (
	"chatgpt-web/pkg/password"

	"gorm.io/gorm"
)

// BeforeSave 保存前
func (user *User) BeforeSave(tx *gorm.DB) (err error) {
	if !password.IsHashed(user.Password) {
		user.Password = password.Hash(user.Password)
	}
	return
}
