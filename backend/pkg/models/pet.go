package models

import "gorm.io/gorm"

type Pet struct {
	gorm.Model
	Name      string
	PhotoUrls string
	Status    string
}

type Tag struct {
	gorm.Model
	Name string
}

type PetTag struct {
	gorm.Model
	PetID uint
	TagID uint
	Pet   Pet `gorm:"foreignKey:PetID;references:ID"`
	Tag   Tag `gorm:"foreignKey:TagID;references:ID"`
}
