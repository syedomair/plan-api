package stat

import (
	"time"

	log "github.com/go-kit/kit/log"
	"github.com/jinzhu/gorm"
	"github.com/syedomair/plan-api/models"
)

type StatRepositoryInterface interface {
	GetTotalUserCount() (string, error)
	GetUserRegData() ([]*models.StatUserRegPerMonth, error)
}

type StatRepository struct {
	Db     *gorm.DB
	Logger log.Logger
}

func (repo *StatRepository) GetTotalUserCount() (string, error) {

	start := time.Now()
	repo.Logger.Log("METHOD", "GetTotalUserCount", "SPOT", "method start", "time_start", start)
	count := "0"
	if err := repo.Db.Table("users").
		Select("*").
		Count(&count).Error; err != nil {
		return "", err
	}

	repo.Logger.Log("METHOD", "GetTotalUserCount", "SPOT", "method end", "time_spent", time.Since(start))

	return count, nil
}

func (repo *StatRepository) GetUserRegData() ([]*models.StatUserRegPerMonth, error) {

	start := time.Now()
	repo.Logger.Log("METHOD", "GetUserRegData", "SPOT", "method start", "time_start", start)
	var userRegs []*models.StatUserRegPerMonth

	if err := repo.Db.Raw("select extract(year from created_at) as year, to_char(created_at, 'mm') as month, count(*) as count from users group by 1, 2 order by 1,2").Scan(&userRegs).Error; err != nil {
		return nil, err
	}

	repo.Logger.Log("METHOD", "GetUserRegData", "SPOT", "method end", "time_spent", time.Since(start))

	return userRegs, nil
}
