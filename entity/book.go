package entity

import "gorm.io/gorm"

type Books struct {
	gorm.Model
	Title string  `valid:"runelength(3|100)~Name of the book must contain 3-100 characters"`
	Price float64 `valid:"range(50|5000)~Price must between 50 and 5000"`
	Code  string  `valid:"matches(^[B][K][0-9]{6}$)"`
}
