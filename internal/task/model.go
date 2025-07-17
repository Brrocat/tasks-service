package task

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Title  string `json:"title"`
	UserID uint   `json:"user_id" gorm:"index"`
}
