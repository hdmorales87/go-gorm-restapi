package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Title string `json:"title" gorm:"type:varchar(100);not null;unique_index"`
	Description string `json:"description"`
	Done bool `json:"done" gorm:"default:false"`
	UserID uint `json:"user_id"`	
}
