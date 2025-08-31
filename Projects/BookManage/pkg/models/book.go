package models

import (
	_ "github.com/Rishi2600/Go-tutorial/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
}
