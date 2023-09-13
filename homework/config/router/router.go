package router

import (
	"homework/app/controllers/studentController"
	"homework/app/controllers/teacherController"
	"homework/app/controllers/userController"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine){
	const pre = "/api"

	api:=r.Group(pre)
	{
		user:=api.Group("/user")
		{
			user.POST("/login",userController.Login)
			user.POST("/reg",userController.Register)
		}
		teacher:=api.Group("/teacher/course")
		{
			teacher.GET("",teacherController.GetClassList)
			teacher.POST("",teacherController.CreateClass)
			teacher.PUT("",teacherController.UpdateClass)
			teacher.DELETE("",teacherController.DeleteClass)
		}
		student:=api.Group("/student")
		{
			student.POST("/course",studentController.CreateStudent)
			student.DELETE("/course",studentController.DeleteStudent)
			student.GET("/select-course",studentController.GetStudentList)
			student.GET("/optional-course",studentController.GetOptionList)
		}
	}
}