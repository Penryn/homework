package teacherController

import (
	"homework/app/models"
	"homework/app/services/teacherService"
	"homework/app/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CreateTeacherDate struct {
	ClassName string `json:"class_name" binding:"required"`
	Time      int    `json:"time" binding:"required"`
	Weekday   int    `json:"weekday" binding:"required"`
	UserID    int    `json:"user_id" binding:"required"`
	Type      int    `json:"type" binding:"required"`
	Number    int    `json:"number" binding:"required"`
}

// 创建课程
func CreateClass(c *gin.Context) {
	var data CreateTeacherDate
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utils.JsonErrorResponse(c, 400, "参数错误")
		return
	}

	var teacher *models.User
	teacher ,err=teacherService.GetNameByUserID(data.UserID)
	if err!=nil{
		utils.JsonInternalServerErrorResponse(c)
		return
	}

    if teacher.Type!=2{
		utils.JsonErrorResponse(c,400,"请回到学生页面")
		return
	}

	err = teacherService.CreateClass(models.Class{
		ClassName: data.ClassName,
		Time:      data.Time,
		Weekday:   data.Weekday,
		UserID:    data.UserID,
		Type:      data.Type,
		Number:    data.Number,
		TeacherName: teacher.Name,
	})
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	utils.JsonSuccessResponse(c, nil)
}

type UpdateTeacherDate struct {
	ClassName string `json:"class_name"`
	Time      int    `json:"time" `
	Weekday   int    `json:"weekday"`
	UserID    int    `json:"user_id" binding:"required"`
	ClassID   int    `json:"class_id" binding:"required"`
	Type      int    `json:"type"`
	Number    int    `json:"number"`
}

// 更新课程信息
func UpdateClass(c *gin.Context) {
	var data UpdateTeacherDate
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utils.JsonErrorResponse(c, 400, "参数错误")
		return
	}
	//获取课程信息
	var class *models.Class
	class, err = teacherService.GetUserIDByClass(data.ClassID)
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	//判断老师是否同一人
	flag := teacherService.CompareUserId(data.UserID, class.UserID)
	if !flag {
		utils.JsonErrorResponse(c, 400, "user不符合")
		return
	}

	err = teacherService.UpdateClassClass(models.Class{
		ClassName: data.ClassName,
		Time:      data.Time,
		Weekday:   data.Weekday,
		UserID:    data.UserID,
		Type:      data.Type,
		Number:    data.Number,
		ClassID:   data.ClassID,
	})
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	utils.JsonSuccessResponse(c, nil)
}

type DeleteTeacherData struct {
	UserID  int `json:"user_id" binding:"required"`
	ClassID int `json:"class_id" binding:"required"`
}

// 删除课程
func DeleteClass(c *gin.Context) {
	var data DeleteTeacherData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utils.JsonErrorResponse(c, 400, "参数错误")
		return
	}
	//获取课程信息
	var class *models.Class
	class, err = teacherService.GetUserIDByClass(data.ClassID)
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}

	//判断增删课程老师是否同一人
	flag := teacherService.CompareUserId(data.UserID, class.UserID)
	if !flag {
		utils.JsonErrorResponse(c, 400, "user不符合")
		return
	}
	if class.Total==0{
		utils.JsonErrorResponse(c,400,"还有学生选课，请勿删除课程")
	}
	//删除课程
	err = teacherService.DeleteClass(data.ClassID)
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	utils.JsonSuccessResponse(c, nil)
}

type GetClassData struct {
	UserID int `form:"user_id" binding:"required"`
}
//获取课程信息
func GetClassList(c *gin.Context) {
	var data GetClassData
	err := c.ShouldBindQuery(&data)
	if err != nil {
		utils.JsonErrorResponse(c, 400, "参数错误")
		return
	}
	var classList []models.Class
	classList, err = teacherService.GetClassList(data.UserID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.JsonErrorResponse(c, 404, "课程为空")
			return
		} else {
			utils.JsonInternalServerErrorResponse(c)
			return
		}
	}
	utils.JsonSuccessResponse(c, gin.H{
		"class_list": classList,
	})
}
