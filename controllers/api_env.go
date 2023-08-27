package controllers

import "gorm.io/gorm"

type APIEnv struct {
	DB *gorm.DB
}
