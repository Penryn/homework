package models



type Class struct {
	ClassID   int      `json:"class_id" gorm:"primaryKey"`
	ClassName string    `json:"class_name"`
	Time      int       `json:"time"`
	Weekday   int       `json:"weekday"`
	UserID    int       `json:"user_id"`
	Type      int       `json:"type"`
	Number    int       `json:"number"`
	Total     int       `json:"total"`
	TeacherName string `json:"teacher_name,omitempty"`

}
