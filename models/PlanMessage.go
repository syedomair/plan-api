package models

type PlanMessage struct {
	Id        string `json:"id" gorm:"column:id"`
	PlanId    string `json:"plan_id" gorm:"plan_id"`
	Message   string `json:"message" gorm:"column:message"`
	Action    string `json:"action" gorm:"column:action"`
	CreatedAt string `json:"created_at" gorm:"column:created_at"`
	UpdatedAt string `json:"updated_at" gorm:"column:updated_at"`
}

func (PlanMessage) TableName() string {
	return "plan_messages"
}
