package teacherService

import (
	"homework/app/models"
	"homework/config/database"
)

func CreateClass(class models.Class) error {
	result := database.DB.Create(&class)
	return result.Error
}

func UpdateClassClass(class models.Class) error {
	//result := database.DB.Omit("user_id").Save(&class)
	result :=database.DB.Model(&class).Updates(models.Class{ClassName: class.ClassName, Time: class.Time, Weekday: class.Weekday,Type: class.Type,Number: class.Number})
	return result.Error
}

func DeleteClass(class_id int) error {
	result := database.DB.Where("class_id=?", class_id).Delete(&models.Class{})
	return result.Error
}

func CompareUserId(u1 int, u2 int) bool {
	return u1 == u2
}



func GetUserIDByClass(cLassid int) (*models.Class, error) {
	var claSS models.Class
	result := database.DB.Where("class_id = ?", cLassid).First(&claSS)
	if result.Error != nil {
		return nil, result.Error
	}
	return &claSS, nil
}

func GetClassList(userID int) ([]models.Class, error) {
	result := database.DB.Where("user_id=?", userID).Find(&models.Class{})
	if result.Error != nil {
		return nil, result.Error
	}
	var classList []models.Class
	result = database.DB.Omit("teacher_name").Where("user_id=?", userID).Find(&classList)
	if result.Error != nil {
		return nil, result.Error
	}
	return classList, nil
}

func GetNameByUserID(userid int) (*models.User, error) {
	var user models.User
	result := database.DB.Where("user_id = ?", userid).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}