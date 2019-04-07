package user

import (
	"encoding/json"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/syedomair/plan-api/models"

	log "github.com/go-kit/kit/log"
	"github.com/jinzhu/gorm"
)

type UserRepositoryInterface interface {
	GetAll(limit string, offset string, orderby string, sort string) ([]*models.UserReduced, string, error)
	GetAllBatch(limit string, offset string, orderby string, sort string) (string, error)
	GetBatchTask(batchTaskId string) (*models.BatchTask, error)
}

type UserRepository struct {
	Db     *gorm.DB
	Logger log.Logger
}

func (repo *UserRepository) GetBatchTask(batchTaskId string) (*models.BatchTask, error) {
	start := time.Now()
	repo.Logger.Log("METHOD", "GetBatchTask", "SPOT", "method start", "time_start", start)
	batchTask := models.BatchTask{}
	if err := repo.Db.Table("batch_tasks").Where("id = ?", batchTaskId).Find(&batchTask).Error; err != nil {
		return nil, err
	}

	repo.Logger.Log("METHOD", "GetBatchTask", "SPOT", "method end", "time_spent", time.Since(start))
	return &batchTask, nil
}
func (repo *UserRepository) GetAll(limit string, offset string, orderby string, sort string) ([]*models.UserReduced, string, error) {

	start := time.Now()
	repo.Logger.Log("METHOD", "GetAll", "SPOT", "method start", "time_start", start)
	var users []*models.UserReduced
	count := "0"
	if err := repo.Db.Table("users").
		Select("*").
		Limit(limit).
		Offset(offset).
		Order(orderby + " " + sort).
		Count(&count).
		Scan(&users).Error; err != nil {
		return nil, "", err
	}

	repo.Logger.Log("METHOD", "GetAll", "SPOT", "method end", "time_spent", time.Since(start))
	return users, count, nil
}

func (repo *UserRepository) GetAllBatch(limit string, offset string, orderby string, sort string) (string, error) {

	start := time.Now()
	repo.Logger.Log("METHOD", "GetAllBatch", "SPOT", "method start", "time_start", start)

	id, _ := uuid.NewV4()
	batchTaskId := id.String()
	newBatchTask := &models.BatchTask{
		Id:          batchTaskId,
		ApiName:     "user_get_all_batch",
		Status:      0,
		CompletedAt: "1900-01-01 00:00:00",
		CreatedAt:   time.Now().Format(time.RFC3339)}

	if err := repo.Db.Create(newBatchTask).Error; err != nil {
		return "", err
	}

	chBatchTaskId := make(chan string)
	go func(batchTaskId string) { chBatchTaskId <- batchTaskId }(batchTaskId)
	go func(batchTaskId string) {
		//to simulate large processing by putting timer for 5 sec
		time.Sleep(5 * time.Second)
		users, _, _ := repo.GetAll(limit, offset, orderby, sort)
		responseDataJson, _ := json.Marshal(users)
		repo.Db.Table("batch_tasks").Where("id = ?", batchTaskId).Updates(map[string]interface{}{"status": 1, "data": responseDataJson, "completed_at": time.Now().Format(time.RFC3339)})
	}(<-chBatchTaskId)

	repo.Logger.Log("METHOD", "GetAllBatch", "SPOT", "method end", "time_spent", time.Since(start))
	return batchTaskId, nil
}
