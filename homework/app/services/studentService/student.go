package studentService

import (
	"homework/app/models"
	"homework/config/database"
)

func GetUserByUserID(userid int) (*models.User, error) {
	var user models.User
	result := database.DB.Where("user_id = ?", userid).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func GetClassByClassID(classid int) (*models.Class, error) {
	var class models.Class
	result := database.DB.Where("class_id = ?", classid).First(&class)
	if result.Error != nil {
		return nil, result.Error
	}
	return &class, nil
}


func CreateClass(student models.Student) error {
	result := database.DB.Create(&student)
	Updatetotal(student.ClassID)
	return result.Error
}

func Updatetotal(uid int) error {
	num := GetTotal(uid)
	//result := database.DB.Omit("user_id").Save(&class)
	database.DB.Model(models.Class{}).Where("class_id=?", uid).Select("total").Updates(models.Class{Total: num})
	return nil
}

func GetTotal(classID int) int {
	result, _ := GetStudentNum(classID)
	num := len(result)
	return num
}

func GetStudentNum(classID int) ([]models.Student, error) {
	result := database.DB.Where("class_id=?", classID).Find(&models.Student{})
	if result.Error != nil {
		return nil, result.Error
	}
	var studentList1 []models.Student
	result = database.DB.Where("class_id=?", classID).Find(&studentList1)
	if result.Error != nil {
		return nil, result.Error
	}
	return studentList1, nil
}

func GetStudentByClassID(cLassid int) (*models.Student, error) {
	var cid models.Student
	result := database.DB.Where("class_id = ?", cLassid).Find(&cid)
	if result.Error != nil {
		return nil, result.Error
	}
	return &cid, nil
}

func CompareUserId(u1 int, u2 int) bool {
	return u1 == u2
}

func DeleteClass(class_id, user_id int) error {
	result := database.DB.Where("user_id=?", user_id).Where("class_id=?", class_id).Delete(&models.Student{})
	Updatetotal(class_id)
	return result.Error
}

func GetClassByClassid(classid int) ([]models.Class, error) {
	var sid []models.Class
	result := database.DB.Where("class_id = ?", classid).Find(&sid)
	if result.Error != nil {
		return nil, result.Error
	}
	return sid, nil
}



func GetClassByUserID(userid int) ([]models.Student, error) {
	var sid []models.Student
	result := database.DB.Where("user_id = ?", userid).Find(&sid)
	if result.Error != nil {
		return nil, result.Error
	}
	return sid, nil
}

func GetClassList(cLassID []int)([]models.Class,error){
	var classList []models.Class
	result := database.DB.Not(cLassID).Where("total<number").Find(&classList)
	if result.Error != nil {
		return nil, result.Error
	}
	return classList, nil

}

func CheckUserClassExist(userid,classid int) error {
	result := database.DB.Where("user_id = ? And class_id = ?" ,userid,classid).First(&models.Student{})
	return result.Error
}