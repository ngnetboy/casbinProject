package model

import "github.com/jinzhu/gorm"

type Account struct {
	gorm.Model
	Name     string `gorm:"type:varchar(64) unique" json:"username"`
	Password string `gorm:"type:varchar(255) not null" json:"password"`
	Role     string `gorm:"type:varchar(64)" json:"role"`
}
