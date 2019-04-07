package models

type BatchTask struct {
	Id          string `json:"id" gorm:"column:id"`
	ApiName     string `json:"api_name" gorm:"column:api_name"`
	Data        string `json:"data" gorm:"column:data"`
	Status      int    `json:"status" gorm:"column:status"`
	CreatedAt   string `json:"created_at" gorm:"column:created_at"`
	CompletedAt string `json:"completed_at" gorm:"column:completed_at"`
}

func (BatchTask) TableName() string {
	return "batch_tasks"
}
