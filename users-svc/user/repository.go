package user

import (
	"time"

	"github.com/syedomair/plan-api/models"

	log "github.com/go-kit/kit/log"
	"github.com/jinzhu/gorm"
)

type UserRepositoryInterface interface {
	GetAll(limit string, offset string, orderby string, sort string) ([]*models.UserReduced, string, error)
}

type UserRepository struct {
	Db     *gorm.DB
	Logger log.Logger
}

func (repo *UserRepository) GetAll(limit string, offset string, orderby string, sort string) ([]*models.UserReduced, string, error) {

	start := time.Now()
	repo.Logger.Log("METHOD", "GetAll", "SPOT", "method start", "time_start", start)
	var users []*models.UserReduced
	count := "0"
	if err := repo.Db.Table("users").
		Select("*").
		Count(&count).
		Limit(limit).
		Offset(offset).
		Order(orderby + " " + sort).
		Scan(&users).Error; err != nil {
		return nil, "", err
	}

	repo.Logger.Log("METHOD", "GetAll", "SPOT", "method end", "time_spent", time.Since(start))
	return users, count, nil
}
