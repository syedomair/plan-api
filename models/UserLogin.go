package models

type UsersLogin struct {
	UserId    string `json:"user_id" gorm:"column:user_id"`
	Token     string `json:"token" gorm:"token"`
	CreatedAt string `json:"created_at" gorm:"column:created_at"`
}

func (UsersLogin) TableName() string {
	return "users_logins"
}
