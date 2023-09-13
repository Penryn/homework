package studentController

import (
	"homework/app/models"
	"homework/app/services/studentService"
	"homework/app/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type GetStudentDaTe struct {
	UserID int `form:"user_id" binding:"required"`
}

func GetOptionList(c *gin.Context){
	var data GetStudentDaTe
	err :=c.ShouldBindQuery(&data)
	if err !=nil{
		utils.JsonErrorResponse(c,400,"参数错误")
		return
	}
	
	//获取用户信息
	var student *models.User
	student, err = studentService.GetUserByUserID(data.UserID)
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	if student.Type!=1{
		utils.JsonErrorResponse(c,400,"请回到教师页面")
		return
	}

	//获取课程信息
	var sclass []models.Student
	sclass,err =studentService.GetClassByUserID(data.UserID)
	if err !=nil{
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	var n []int
	for _,i :=range sclass{
		n=append(n, i.ClassID)
	}
	var classList []models.Class
	classList,err=studentService.GetClassList(n)
	if err!=nil{
		if err ==gorm.ErrRecordNotFound{
			utils.JsonErrorResponse(c,404,"课程为空")
			return
		}else{
			utils.JsonInternalServerErrorResponse(c)
			return
		}
	}
	utils.JsonSuccessResponse(c,gin.H{
		"class_list":classList,
	})
}