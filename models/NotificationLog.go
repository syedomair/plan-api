package models

type NotificationLog struct {
	Id             string `json:"id" gorm:"column:id"`
	NotificationId string `json:"notification_id" gorm:"column:notification_id"`
	Error          string `json:"error" gorm:"column:error"`
	retriedAt      string `json:"retried_at" gorm:"column:retried_at"`
	CreatedAt      string `json:"created_at" gorm:"column:created_at"`
}

func (NotificationLog) TableName() string {
	return "notification_logs"
}
