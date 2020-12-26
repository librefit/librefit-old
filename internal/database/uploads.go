package database

import "gorm.io/gorm"

type Upload struct {
	gorm.Model
	Title       string
	Caption     string
	Size        string
	RelativeDir string
	Filename    string
	UserID      uint
}
