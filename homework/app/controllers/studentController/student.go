package studentController

import (
	"homework/app/models"
	"homework/app/services/studentService"
	"homework/app/utils"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CreateStudentDate struct {
	UserID  int `json:"user_id" binding:"required"`
	ClassID int `json:"class_id" binding:"required"`
}

// 学生选课
func CreateStudent(c *gin.Context) {
	var data CreateStudentDate
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utils.JsonErrorResponse(c, 400, "参数错误")
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
	var class *models.Class
	class,err =studentService.GetClassByClassID(data.ClassID)
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	if class.Total>=class.Number{
		utils.JsonErrorResponse(c,400,"课程已满")
		return
	}
	//检测是否重新报名
	err =studentService.CheckUserClassExist(data.UserID,data.ClassID)
	if err==nil{
		if err !=gorm.ErrRecordNotFound{
			utils.JsonErrorResponse(c,401,"请不要重新报名")
		}else{
			utils.JsonInternalServerErrorResponse(c)
		}
		return
	}

	time1 := time.Now().Format("2006-01-02")
	err = studentService.CreateClass(models.Student{
		ClassID:     data.ClassID,
		UserID:      uint(data.UserID),
		Date:        time1,
	})
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	utils.JsonSuccessResponse(c, nil)
}



type DeleteStudentData struct {
	UserID  int `json:"user_id" binding:"required"`
	ClassID int `json:"class_id" binding:"required"`
}

// 学生退课
func DeleteStudent(c *gin.Context) {
	var data DeleteStudentData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utils.JsonErrorResponse(c, 400, "参数错误")
		return
	}
	//获取课程信息
	var student *models.Student
	student, err = studentService.GetStudentByClassID(data.ClassID)
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	//判断选退课程是否同一人
	flag := studentService.CompareUserId(data.UserID, int(student.UserID))
	if !flag {
		utils.JsonErrorResponse(c, 400, "user不符合")
		return
	}
	//退课
	err = studentService.DeleteClass(data.ClassID, data.UserID)
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	utils.JsonSuccessResponse(c, nil)
}
