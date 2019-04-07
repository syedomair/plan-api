package models

type User struct {
	Id        string `json:"id" gorm:"column:id"`
	Email     string `json:"email" gorm:"column:email"`
	FirstName string `json:"first_name" gorm:"column:first_name"`
	LastName  string `json:"last_name" gorm:"column:last_name"`
	Password  string `json:"password" gorm:"column:password"`
	Verified  string `json:"verified" gorm:"column:verified"`
	IsAdmin   string `json:"is_admin" gorm:"column:is_admin"`
	CreatedAt string `json:"created_at" gorm:"column:created_at"`
	UpdatedAt string `json:"updated_at" gorm:"column:updated_at"`
}

type UserReduced struct {
	Id        string `json:"id" gorm:"column:id"`
	Email     string `json:"email" gorm:"column:email"`
	FirstName string `json:"first_name" gorm:"column:first_name"`
	LastName  string `json:"last_name" gorm:"column:last_name"`
	IsAdmin   string `json:"is_admin" gorm:"column:is_admin"`
	Verified  string `json:"verified" gorm:"column:verified"`
	CreatedAt string `json:"created_at" gorm:"column:created_at"`
	UpdatedAt string `json:"updated_at" gorm:"column:updated_at"`
}

func (User) TableName() string {
	return "users"
}
