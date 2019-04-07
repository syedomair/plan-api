package models

type Notification struct {
	Id              string `json:"id" gorm:"column:id"`
	NotificationMsg string `json:"notification_msg" gorm:"column:notification_msg"`
	Object          string `json:"object" gorm:"column:object"`
	Operation       string `json:"operation" gorm:"column:operation"`
	CreatedAt       string `json:"created_at" gorm:"column:created_at"`
}

func (Notification) TableName() string {
	return "notifications"
}
