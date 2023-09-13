package studentController

import (
	"homework/app/models"
	"homework/app/services/studentService"
	"homework/app/utils"

	"github.com/gin-gonic/gin"
)

type GetStudentData struct {
	UserID int `form:"user_id" binding:"required"`
}

func GetStudentList(c *gin.Context) {
	var data GetStudentData
	err := c.ShouldBindQuery(&data)
	if err != nil {
		utils.JsonErrorResponse(c, 400, "参数错误")
		return
	}

	type Sctclass struct {
		ClassID     int       `json:"class_id"`
		ClassName   string    `json:"class_name"`
		Time        int       `json:"time"`
		Weekday     int       `json:"weekday"`
		UserID      int       `json:"user_id"`
		Type        int       `json:"type"`
		Number      int       `json:"number"`
		Total       int       `json:"total"`
		TeacherName string    `json:"teacher_name"`
		Date        string `json:"date"`
	}
	var sclass []models.Student
	sclass,err =studentService.GetClassByUserID(data.UserID)
	if err !=nil{
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	var new []Sctclass
	for _,j :=range sclass{
		slass,_:=studentService.GetClassByClassid(j.ClassID)
		for _,i:=range slass{
			classdata:=Sctclass{
				ClassID: j.ClassID,
				UserID: i.UserID,
				TeacherName: i.TeacherName,
				ClassName: i.ClassName,
				Date: j.Date,
				Time: i.Time,
				Weekday: i.Weekday,
				Type: i.Type,
				Number: i.Number,
				Total: i.Total,
			}
			new =append(new, classdata)
		}
		

	}
	utils.JsonSuccessResponse(c,gin.H{
		"class_list":new,
	})

}
