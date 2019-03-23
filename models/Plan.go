package models

type Plan struct {
	Id        string `json:"id" gorm:"column:id"`
	Title     string `json:"title" gorm:"column:title"`
	Status    int    `json:"status" gorm:"column:status"`
	Validity  int    `json:"validity" gorm:"column:validity"`
	Cost      int    `json:"cost" gorm:"column:cost"`
	CreatedAt string `json:"created_at" gorm:"column:created_at"`
	UpdatedAt string `json:"updated_at" gorm:"column:updated_at"`
}

func (Plan) TableName() string {
	return "plans"
}
