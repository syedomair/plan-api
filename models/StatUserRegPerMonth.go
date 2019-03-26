package models

type StatUserRegPerMonth struct {
	Year  string `json:"year" gorm:"column:year"`
	Month string `json:"month" gorm:"column:month"`
	Count string `json:"count" gorm:"column:count"`
}
