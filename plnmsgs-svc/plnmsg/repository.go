package plnmsg

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/syedomair/plan-api/models"

	log "github.com/go-kit/kit/log"
	"github.com/jinzhu/gorm"
)

type PlanMessageRepositoryInterface interface {
	Create(inputPlanMsg map[string]interface{}, planId string) (string, error)
	GetAll(limit string, offset string, orderby string, sort string, planId string) ([]*models.PlanMessage, string, error)
	Get(planMessageId string) (*models.PlanMessage, error)
	Update(inputPlanMsg map[string]interface{}, planMessageId string) error
	Delete(planMessage models.PlanMessage) error
}

type PlanMessageRepository struct {
	Db     *gorm.DB
	Logger log.Logger
}

func (repo *PlanMessageRepository) Create(inputPlanMsg map[string]interface{}, planId string) (string, error) {

	start := time.Now()
	repo.Logger.Log("METHOD", "Create", "SPOT", "method start", "time_start", start)

	message := ""
	if messageValue, ok := inputPlanMsg["message"]; ok {
		message = messageValue.(string)
	}
	action := ""
	if actionValue, ok := inputPlanMsg["action"]; ok {
		action = actionValue.(string)
	}
	id, _ := uuid.NewV4()
	planMessageId := id.String()
	newPlanMessage := &models.PlanMessage{
		Id:        planMessageId,
		PlanId:    planId,
		Message:   message,
		Action:    action,
		CreatedAt: time.Now().Format(time.RFC3339),
		UpdatedAt: time.Now().Format(time.RFC3339)}

	if err := repo.Db.Create(newPlanMessage).Error; err != nil {
		return "", err
	}
	repo.Logger.Log("METHOD", "Create", "SPOT", "method end", "time_spent", time.Since(start))
	return planMessageId, nil
}

func (repo *PlanMessageRepository) GetAll(limit string, offset string, orderby string, sort string, planId string) ([]*models.PlanMessage, string, error) {

	start := time.Now()
	repo.Logger.Log("METHOD", "GetAll", "SPOT", "method start", "time_start", start)
	var planMessages []*models.PlanMessage
	count := "0"
	if err := repo.Db.Table("plan_messages").
		Select("*").
		Count(&count).
		Limit(limit).
		Offset(offset).
		Order(orderby+" "+sort).
		Where("plan_id = ?", planId).
		Scan(&planMessages).Error; err != nil {
		return nil, "", err
	}
	repo.Logger.Log("METHOD", "GetAll", "SPOT", "method end", "time_spent", time.Since(start))
	return planMessages, count, nil
}

func (repo *PlanMessageRepository) Get(planMessageId string) (*models.PlanMessage, error) {
	start := time.Now()
	repo.Logger.Log("METHOD", "Get", "SPOT", "method start", "time_start", start)
	planMessage := models.PlanMessage{}
	if err := repo.Db.Where("id = ?", planMessageId).Find(&planMessage).Error; err != nil {
		return nil, err
	}

	repo.Logger.Log("METHOD", "Get", "SPOT", "method end", "time_spent", time.Since(start))
	return &planMessage, nil
}

func (repo *PlanMessageRepository) Update(inputPlanMsg map[string]interface{}, planMessageId string) error {
	start := time.Now()
	repo.Logger.Log("METHOD", "Update", "SPOT", "method start", "time_start", start)
	if err := repo.Db.Table("plan_messages").Where("id = ?", planMessageId).Updates(inputPlanMsg).Error; err != nil {
		return err
	}
	repo.Logger.Log("METHOD", "Update", "SPOT", "method end", "time_spent", time.Since(start))
	return nil
}

func (repo *PlanMessageRepository) Delete(planMessage models.PlanMessage) error {
	start := time.Now()
	repo.Logger.Log("METHOD", "Delete", "SPOT", "method start", "time_start", start)
	planMessageId := planMessage.Id
	if err := repo.Db.Where("id = ?", planMessageId).Find(&planMessage).Error; err != nil {
		return err
	}
	if err := repo.Db.Delete(&planMessage).Error; err != nil {
		return err
	}
	repo.Logger.Log("METHOD", "Delete", "SPOT", "method end", "time_spent", time.Since(start))
	return nil
}
