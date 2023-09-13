package userService

import (
	"homework/app/models"
	"homework/config/database"
)


func CheckUserExistByAccount(aCCount string) error {
	result := database.DB.Where("account = ?", aCCount).First(&models.User{})
	return result.Error
}

func GetUserByAccount(aCCount string) (*models.User, error) {
	var user models.User
	result := database.DB.Where("account = ?", aCCount).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
func ComparePwd(pwd1 string, pwd2 string) bool {
	return pwd1 == pwd2
}

func Register(user models.User) error {
	result := database.DB.Create(&user)
	return result.Error
}
