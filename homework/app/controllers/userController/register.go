package userController

import (
	"homework/app/models"
	"homework/app/services/userService"
	"homework/app/utils"
	"regexp"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type RegisterData struct {
	Name     string `json:"name" binding:"required"`
	Type     int8   `json:"type" binding:"required"`
	Account  string `json:"account" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// 注册
func Register(c *gin.Context) {
	var data RegisterData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utils.JsonErrorResponse(c, 400, "参数错误")
		return
	}
	if len(data.Password) <= 8 {
		utils.JsonErrorResponse(c, 400, "密码太短")
		return
	}
	sample := regexp.MustCompile(`^[0-9]+$`)
	if !sample.MatchString(data.Account) {
		utils.JsonErrorResponse(c, 400, "账号应为数字")
		return
	}

	// 判断手机号是否已经注册
	err = userService.CheckUserExistByAccount(data.Account)
	if err == nil {
		utils.JsonErrorResponse(c, 400, "手机号已注册")
		return
	} else if err != nil && err != gorm.ErrRecordNotFound {
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	//加密
	pwd,err:=bcrypt.GenerateFromPassword([]byte(data.Password),bcrypt.DefaultCost)
	if err !=nil{
		utils.JsonInternalServerErrorResponse(c)
		return
	}


	// 注册用户
	err = userService.Register(models.User{
		Name:     data.Name,
		Type:     data.Type,
		Account:  data.Account,
		Password: pwd,
	})
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}

	utils.JsonSuccessResponse(c, nil)
}
