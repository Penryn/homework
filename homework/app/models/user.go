package models

type User struct {
	UserID   int   `json:"user_id" gorm:"primaryKey"`
	Name     string `json:"-"`
	Account  string `json:"-"`
	Password []byte `json:"-"`
	Type     int8   `json:"type"`
}
