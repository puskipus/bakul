package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Nama     string `gorm:"varchar(100)" json:"nama"`
	Telepon  string `gorm:"varchar(15)" json:"telepon"`
	Email    string `gorm:"varchar(100)" json:"email"`
	Password string `gorm:"varchar(100)" json:"password"`
	Foto     string `gorm:"varchar" json:"foto"`
}
