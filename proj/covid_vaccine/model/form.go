package model

import "github.com/jinzhu/gorm"

type Details struct {
	gorm.Model
	paitentName string
	age         int
	mnumber     int
}
