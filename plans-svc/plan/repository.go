package plan

import (
	"strconv"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/syedomair/plan-api/models"

	log "github.com/go-kit/kit/log"
	"github.com/jinzhu/gorm"
)

type PlanRepositoryInterface interface {
	Create(inputPlan map[string]interface{}) (string, error)
	GetAll(limit string, offset string, orderby string, sort string) ([]*models.Plan, string, error)
	Get(planId string) (*models.Plan, error)
	Update(inputPlan map[string]interface{}, planId string) error
	Delete(plan models.Plan) error
}

type PlanRepository struct {
	Db     *gorm.DB
	Logger log.Logger
}

func (repo *PlanRepository) Create(inputPlan map[string]interface{}) (string, error) {

	start := time.Now()
	repo.Logger.Log("METHOD", "Create", "SPOT", "method start", "time_start", start)

	title := ""
	if titleValue, ok := inputPlan["title"]; ok {
		title = titleValue.(string)
	}
	status := 0
	if statusValue, ok := inputPlan["status"]; ok {
		status, _ = strconv.Atoi(statusValue.(string))
	}
	cost := 0
	if costValue, ok := inputPlan["cost"]; ok {
		cost, _ = strconv.Atoi(costValue.(string))
	}
	validity := 0
	if validityValue, ok := inputPlan["validity"]; ok {
		validity, _ = strconv.Atoi(validityValue.(string))
	}

	id, _ := uuid.NewV4()
	planId := id.String()
	newPlan := &models.Plan{
		Id:        planId,
		Title:     title,
		Status:    status,
		Cost:      cost,
		Validity:  validity,
		CreatedAt: time.Now().Format(time.RFC3339),
		UpdatedAt: time.Now().Format(time.RFC3339)}

	if err := repo.Db.Create(newPlan).Error; err != nil {
		return "", err
	}

	repo.Logger.Log("METHOD", "Create", "SPOT", "method end", "time_spent", time.Since(start))
	return planId, nil
}

func (repo *PlanRepository) GetAll(limit string, offset string, orderby string, sort string) ([]*models.Plan, string, error) {

	start := time.Now()
	repo.Logger.Log("METHOD", "GetAll", "SPOT", "method start", "time_start", start)
	var plans []*models.Plan
	count := "0"
	if err := repo.Db.Table("plans").
		Select("*").
		Count(&count).
		Limit(limit).
		Offset(offset).
		Order("plans." + orderby + " " + sort).
		Scan(&plans).Error; err != nil {
		return nil, "", err
	}
	repo.Logger.Log("METHOD", "GetAll", "SPOT", "method end", "time_spent", time.Since(start))
	return plans, count, nil
}

func (repo *PlanRepository) Get(planId string) (*models.Plan, error) {
	start := time.Now()
	repo.Logger.Log("METHOD", "Get", "SPOT", "method start", "time_start", start)
	plan := models.Plan{}
	if err := repo.Db.Table("plans").Where("id = ?", planId).Find(&plan).Error; err != nil {
		return nil, err
	}

	repo.Logger.Log("METHOD", "Get", "SPOT", "method end", "time_spent", time.Since(start))
	return &plan, nil
}

func (repo *PlanRepository) Update(inputPlan map[string]interface{}, planId string) error {
	start := time.Now()
	repo.Logger.Log("METHOD", "Update", "SPOT", "method start", "time_start", start)
	if err := repo.Db.Table("plans").Where("id = ?", planId).Updates(inputPlan).Error; err != nil {
		return err
	}
	repo.Logger.Log("METHOD", "Update", "SPOT", "method end", "time_spent", time.Since(start))
	return nil
}

func (repo *PlanRepository) Delete(plan models.Plan) error {
	start := time.Now()
	repo.Logger.Log("METHOD", "Delete", "SPOT", "method start", "time_start", start)
	planId := plan.Id
	if err := repo.Db.Where("id = ?", planId).Find(&plan).Error; err != nil {
		return err
	}
	if err := repo.Db.Delete(&plan).Error; err != nil {
		return err
	}
	repo.Logger.Log("METHOD", "Delete", "SPOT", "method end", "time_spent", time.Since(start))
	return nil
}
