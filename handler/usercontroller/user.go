package user

import (
	"errors"
	"github.com/WebDesign/handler"
	"github.com/WebDesign/log"
	"github.com/WebDesign/model/requests"
	"github.com/WebDesign/model/response"
	"github.com/WebDesign/model/user"
	"github.com/gin-gonic/gin"
)

type UsersController struct {
	handler.BaseAPIController
}

// CurrentUser 当前登录用户信息
func GetCurrentUser(c *gin.Context) {
	userModel := CurrentUser(c)
	response.Data(c, userModel)
}

// CurrentUser 从 gin.context 中获取当前登录用户
func CurrentUser(c *gin.Context) user.User {
	userModel, ok := c.MustGet("current_user").(user.User)
	if !ok {
		log.LogIf(errors.New("无法获取用户"))
		return user.User{}
	}
	// db is now a *DB value
	return userModel
}

// GetUsers所有用户
func GetUsers(c *gin.Context) {
	request := requests.PaginationRequest{}
	if ok := handler.Validate(c, &request, requests.Pagination); !ok {
		return
	}

	data, pager := user.Paginate(c, 10)
	response.JSON(c, gin.H{
		"data":  data,
		"pager": pager,
	})
}

func UpdateProfile(c *gin.Context) {

	request := requests.UserUpdateProfileRequest{}
	if ok := handler.Validate(c, &request, requests.UserUpdateProfile); !ok {
		return
	}

	currentUser := CurrentUser(c)
	currentUser.Name = request.Name
	currentUser.City = request.City
	currentUser.Introduction = request.Introduction
	rowsAffected := currentUser.Save()
	if rowsAffected > 0 {
		response.Data(c, currentUser)
	} else {
		response.Abort500(c, "更新失败，请稍后尝试~")
	}
}
