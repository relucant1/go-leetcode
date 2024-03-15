package models

import (
	"github.com/jinzhu/gorm/"
	"github.com/relucat1/go-bookstore/pkg/config"
)

var db *gorm.DB

type Boook struct {
	gorm.Model
	Name        string `gorm:"name"`
	Author      string `gorm:"author"`
	Description string `gorm:"description"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Boook{})
}
