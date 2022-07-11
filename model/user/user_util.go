package user

//用户列表分页
import (
	db "github.com/WebDesign/database"
	"github.com/WebDesign/pkg/paginator"
	"github.com/gin-gonic/gin"
)

// GetByPhone 通过手机号来获取用户
func GetByPhone(phone string) (userModel User) {
	db.DB.Where("phone = ?", phone).First(&userModel)
	return
}

// GetByMulti 通过 手机号/Email/用户名 来获取用户
func GetByMulti(loginID string) (userModel User) {
	db.DB.
		Where("phone = ?", loginID).
		Or("email = ?", loginID).
		Or("name = ?", loginID).
		First(&userModel)
	return
}

// Get 通过 ID 获取用户
func Get(idstr string) (userModel User) {
	db.DB.Where("id", idstr).First(&userModel)
	return
}

// GetByEmail 通过 Email 来获取用户
func GetByEmail(email string) (userModel User) {
	db.DB.Where("email = ?", email).First(&userModel)
	return
}

// All 获取所有用户数据
func All() (users []User) {
	db.DB.Find(&users)
	return
}

// Paginate 分页内容
func Paginate(c *gin.Context, perPage int) (users []User, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		db.DB.Model(User{}),
		&users,
		"http://localhost:8080",
		perPage,
	)
	return
}
