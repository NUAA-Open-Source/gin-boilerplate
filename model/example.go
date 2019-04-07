package model

import "github.com/jinzhu/gorm"

type Example struct {
	gorm.Model
	Message string
}
