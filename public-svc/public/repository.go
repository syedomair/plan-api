package public

import (
	b64 "encoding/base64"
	"errors"
	"math/rand"
	"time"

	log "github.com/go-kit/kit/log"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"github.com/syedomair/plan-api/models"
)

type PublicRepositoryInterface interface {
	IsEmailUnique(email string) error
	CreateUser(inputUser map[string]interface{}) (*models.User, error)
	CreateUserLogin(userId string, token string) error
	ValidateEmailPasswordFromDB(email string, password string) (*models.User, error)
	FindToken(token string) (models.UsersLogin, error)
}

type PublicRepository struct {
	Db     *gorm.DB
	Logger log.Logger
}

func (repo *PublicRepository) IsEmailUnique(email string) error {
	repo.Logger.Log("METHOD", "IsEmailUnique", "SPOT", "method start")
	start := time.Now()

	user := models.User{}
	if err := repo.Db.Where("email = ?", email).Find(&user).Error; err == nil {
		return errors.New("User email already exists.")
	}
	repo.Logger.Log("METHOD", "IsEmailUnique", "SPOT", "METHOD END", "time_spent", time.Since(start))
	return nil
}

func (repo *PublicRepository) CreateUser(inputUser map[string]interface{}) (*models.User, error) {
	repo.Logger.Log("METHOD", "CreateUser", "SPOT", "method start")
	start := time.Now()

	id, _ := uuid.NewV4()
	userId := id.String()
	password := b64.StdEncoding.EncodeToString([]byte(inputUser["password"].(string)))
	user := models.User{
		Id:        userId,
		FirstName: inputUser["first_name"].(string),
		LastName:  inputUser["last_name"].(string),
		Email:     inputUser["email"].(string),
		Password:  password,
		Verified:  "0",
		IsAdmin:   "0",
		CreatedAt: time.Now().Format(time.RFC3339),
		UpdatedAt: time.Now().Format(time.RFC3339)}

	repo.Db.NewRecord(user)
	if err := repo.Db.Create(&user).Error; err != nil {
		return &user, err
	}

	repo.Logger.Log("METHOD", "CreateUser", "SPOT", "METHOD END", "time_spent", time.Since(start))
	return &user, nil
}

func (repo *PublicRepository) CreateUserLogin(userId string, token string) error {
	repo.Logger.Log("METHOD", "CreateUserLogin", "SPOT", "method start")
	start := time.Now()

	usersLogin := models.UsersLogin{
		UserId:    userId,
		Token:     token,
		CreatedAt: time.Now().Format(time.RFC3339)}

	if err := repo.Db.Where("user_id = ?", userId).Find(&usersLogin).Error; err != nil {
		repo.Db.NewRecord(usersLogin)
		if err := repo.Db.Create(&usersLogin).Error; err != nil {
			return err
		}

	} else {
		if err := repo.Db.Model(&usersLogin).Where("user_id = ?", userId).Updates(models.UsersLogin{Token: token, CreatedAt: time.Now().Format(time.RFC3339)}).Error; err != nil {
			return err
		}
	}
	repo.Logger.Log("METHOD", "CreateUserLogin", "SPOT", "METHOD END", "time_spent", time.Since(start))
	return nil
}

func (repo *PublicRepository) ValidateEmailPasswordFromDB(email string, password string) (*models.User, error) {

	repo.Logger.Log("METHOD", "ValidateEmailPasswordFromDB", "SPOT", "method start")
	start := time.Now()
	user := models.User{}
	encPassword := b64.StdEncoding.EncodeToString([]byte(password))
	if err := repo.Db.Where("email = ?", email).Where("password = ?", encPassword).Find(&user).Error; err != nil {
		return nil, err
	}
	repo.Logger.Log("METHOD", "ValidateEmailPasswordFromDB", "SPOT", "METHOD END", "time_spent", time.Since(start))
	return &user, nil
}

func (repo *PublicRepository) FindToken(token string) (models.UsersLogin, error) {

	repo.Logger.Log("METHOD", "FindToken", "SPOT", "METHOD START")
	start := time.Now()
	usersLogin := models.UsersLogin{}
	if err := repo.Db.Where("token = ?", token).Find(&usersLogin).Error; err != nil {
		return usersLogin, err
	}
	repo.Logger.Log("METHOD", "FindToken", "SPOT", "METHOD END", "time_spent", time.Since(start))
	return usersLogin, nil
}

func (repo *PublicRepository) GenerateApiKey() string {
	var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	b := make([]rune, 35)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}
