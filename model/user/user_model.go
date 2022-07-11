package user

import (
	db "github.com/WebDesign/database"
	"github.com/WebDesign/model"
)

type User struct {
	model.BaseModel
	Name         string `json:"name,omitempty"`
	City         string `json:"city,omitempty"`
	Introduction string `json:"introduction,omitempty"`
	Avatar       string `json:"avatar,omitempty"`
	Email        string `json:"-"`
	Phone        string `json:"-"`
	Password     string `json:"-"`
	model.CommonTimestampsField
}

// Create 创建用户，通过 User.ID 来判断是否创建成功
func (userModel *User) Create() {
	db.DB.Create(&userModel)
}

// ComparePassword 密码是否正确
func (userModel *User) ComparePassword(password string) bool {
	if password == userModel.Password {
		return true
	} else {
		return false
	}
}

func (userModel *User) Save() (rowsAffected int64) {
	result := db.DB.Save(&userModel)
	return result.RowsAffected
}
