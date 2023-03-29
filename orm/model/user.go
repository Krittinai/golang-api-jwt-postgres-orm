package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username  string
	Password  string
	Firstname string
	Lastname  string
}
