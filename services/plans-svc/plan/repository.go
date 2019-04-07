package plan

import (
	"encoding/json"
	"strconv"
	"time"

	uuid "github.com/satori/go.uuid"
	lib "github.com/syedomair/plan-api/lib"
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
	incrementPlanCount()
	decrementPlanCount()
	postPlanNotification(planId string, operation string)
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

	chPlanId := make(chan string)
	go func(planId string) { chPlanId <- planId }(planId)
	go func(planId string) {
		repo.incrementPlanCount()
		repo.postPlanNotification(planId, "Create")
	}(<-chPlanId)
	repo.Logger.Log("METHOD", "Create", "SPOT", "method end", "time_spent", time.Since(start))
	return planId, nil
}

func (repo *PlanRepository) incrementPlanCount() {
	repo.Logger.Log("METHOD", "incrementPlanCount", "SPOT", "METHOD START")
	start := time.Now()

	type Result struct {
		Count string
	}
	var result Result
	if err := repo.Db.Raw("select total_plan as count from stat ").Scan(&result).Error; err != nil {
		return
	}

	planCount, _ := strconv.Atoi(result.Count)

	if err := repo.Db.Table("stat").Updates(map[string]interface{}{"total_plan": planCount + 1}).Error; err != nil {
		return
	}

	repo.Logger.Log("METHOD", "incrementPlanCount", "SPOT", "METHOD END", "time_spent", time.Since(start))
	return
}
func (repo *PlanRepository) decrementPlanCount() {
	repo.Logger.Log("METHOD", "decrementPlanCount", "SPOT", "METHOD START")
	start := time.Now()

	type Result struct {
		Count string
	}
	var result Result
	if err := repo.Db.Raw("select total_plan as count from stat ").Scan(&result).Error; err != nil {
		return
	}

	planCount, _ := strconv.Atoi(result.Count)

	if err := repo.Db.Table("stat").Updates(map[string]interface{}{"total_plan": planCount - 1}).Error; err != nil {
		return
	}

	repo.Logger.Log("METHOD", "decrementPlanCount", "SPOT", "METHOD END", "time_spent", time.Since(start))
	return
}

func (repo *PlanRepository) postPlanNotification(planId string, operation string) {
	repo.Logger.Log("METHOD", "postPlanNotification", "SPOT", "METHOD START")
	start := time.Now()

	plan := models.Plan{}
	if err := repo.Db.Table("plans").Where("id = ?", planId).Find(&plan).Error; err != nil {
		return
	}

	planJson := map[string]string{
		"id":       plan.Id,
		"title":    plan.Title,
		"status":   strconv.Itoa(plan.Status),
		"validity": strconv.Itoa(plan.Validity),
		"cost":     strconv.Itoa(plan.Cost),
	}

	planJsonStr, _ := json.Marshal(planJson)
	notiId, _ := repo.createNotification(string(planJsonStr), "Plan", operation)

	resultCh := make(chan string)
	ticker := time.NewTicker(10 * time.Second)
	go func() {
		for _ = range ticker.C {
			go lib.PostNotificationToHttpBin("post", planJson, resultCh)
		}
	}()

	select {
	case result := <-resultCh:
		if result == "success" {
			_, _ = repo.createNotificationLog(notiId, "")
			ticker.Stop()
		}
	case <-time.After(55 * time.Second):
		_, _ = repo.createNotificationLog(notiId, "Could not Post the message... Timedout. ")
		ticker.Stop()
	}

	repo.Logger.Log("METHOD", "postPlanNotification", "SPOT", "METHOD END", "time_spent", time.Since(start))
	return
}
func (repo *PlanRepository) GetAll(limit string, offset string, orderby string, sort string) ([]*models.Plan, string, error) {

	start := time.Now()
	repo.Logger.Log("METHOD", "GetAll", "SPOT", "method start", "time_start", start)
	var plans []*models.Plan
	count := "0"
	if err := repo.Db.Table("plans").
		Select("*").
		Limit(limit).
		Offset(offset).
		Order("plans." + orderby + " " + sort).
		Count(&count).
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

	go repo.postPlanNotification(planId, "Update")

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

	go repo.decrementPlanCount()

	repo.Logger.Log("METHOD", "Delete", "SPOT", "method end", "time_spent", time.Since(start))
	return nil
}

func (repo *PlanRepository) createNotification(notificationMsg string, object string, operation string) (string, error) {

	start := time.Now()
	repo.Logger.Log("METHOD", "createNotification", "SPOT", "method start", "time_start", start)

	id, _ := uuid.NewV4()
	notificationId := id.String()
	newNotification := &models.Notification{
		Id:              notificationId,
		NotificationMsg: notificationMsg,
		Object:          object,
		Operation:       operation,
		CreatedAt:       time.Now().Format(time.RFC3339)}

	if err := repo.Db.Create(newNotification).Error; err != nil {
		return "", err
	}

	repo.Logger.Log("METHOD", "createNotification", "SPOT", "method end", "time_spent", time.Since(start))
	return notificationId, nil
}
func (repo *PlanRepository) createNotificationLog(notificationId string, errorStr string) (string, error) {

	start := time.Now()
	repo.Logger.Log("METHOD", "createNotificationLog", "SPOT", "method start", "time_start", start)

	id, _ := uuid.NewV4()
	notificationLogId := id.String()
	newNotificationLog := &models.NotificationLog{
		Id:             notificationLogId,
		NotificationId: notificationId,
		Error:          errorStr,
		CreatedAt:      time.Now().Format(time.RFC3339)}

	if err := repo.Db.Create(newNotificationLog).Error; err != nil {
		return "", err
	}

	repo.Logger.Log("METHOD", "createNotificationLog", "SPOT", "method end", "time_spent", time.Since(start))
	return notificationLogId, nil
}
