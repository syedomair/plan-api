package stat

import (
	"time"

	log "github.com/go-kit/kit/log"
	"github.com/jinzhu/gorm"
	"github.com/syedomair/plan-api/models"
)

type StatRepositoryInterface interface {
	GetTotalUserCount() (string, error)
	GetTotalUserCountLast30Days() (string, error)
	GetUserRegData() ([]*models.StatUserRegPerMonth, error)
	GetPlanData() (string, error)
}

type StatRepository struct {
	Db     *gorm.DB
	Logger log.Logger
}

func (repo *StatRepository) GetTotalUserCountLast30Days() (string, error) {

	start := time.Now()
	repo.Logger.Log("METHOD", "GetTotalUserCountLast30Days", "SPOT", "method start", "time_start", start)
	type Result struct {
		Count string
	}
	var result Result
	if err := repo.Db.Raw("select count(*) as count from users where created_at >= NOW() - interval '30 day'").Scan(&result).Error; err != nil {
		return "", err
	}
	repo.Logger.Log("METHOD", "GetTotalUserCountLast30Days", "SPOT", "method end", "time_spent", time.Since(start))
	return result.Count, nil
}
func (repo *StatRepository) GetTotalUserCount() (string, error) {

	start := time.Now()
	repo.Logger.Log("METHOD", "GetTotalUserCount", "SPOT", "method start", "time_start", start)
	type Result struct {
		Count string
	}
	var result Result
	if err := repo.Db.Raw("select total_user as count from stat ").Scan(&result).Error; err != nil {
		return "", err
	}

	repo.Logger.Log("METHOD", "GetTotalUserCount", "SPOT", "method end", "time_spent", time.Since(start))
	return result.Count, nil
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

func (repo *StatRepository) GetPlanData() (string, error) {

	start := time.Now()
	repo.Logger.Log("METHOD", "GetPlanData", "SPOT", "method start", "time_start", start)

	type Result struct {
		Count string
	}
	var result Result
	if err := repo.Db.Raw("select total_plan as count from stat ").Scan(&result).Error; err != nil {
		return "", err
	}
	repo.Logger.Log("METHOD", "GetPlanData", "SPOT", "method end", "time_spent", time.Since(start))

	return result.Count, nil
}
