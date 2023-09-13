package models


type Student struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	UserID      uint    `json:"user_id"`
	ClassID     int    `json:"class_id"`
	Date        string `json:"date" `
}
