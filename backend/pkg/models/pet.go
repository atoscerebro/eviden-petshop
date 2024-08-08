package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Pet struct {
	gorm.Model
	Name      string         `gorm:"not null" binding:"required"`
	PhotoUrls pq.StringArray `gorm:"type:text[];not null" binding:"required"`
	Tags      []Tag          `gorm:"many2many:pet_tags;"`
	Status    string
}

type Tag struct {
	gorm.Model
	Name string
	Pets []Pet `gorm:"many2many:pet_tags;"`
}

// TODO Better representation of TAGS
// type PetTag struct {
// 	gorm.Model
// 	PetID uint
// 	TagID uint
// 	Pet   Pet `gorm:"foreignKey:PetID;references:ID"`
// 	Tag   Tag `gorm:"foreignKey:TagID;references:ID"`
// }
