package user

import (
	"testing"
	"time"

	"github.com/syedomair/plan-api/lib"
)

func TestUserDB(t *testing.T) {

	db, _ := lib.CreateDBConnection()
	repo := &UserRepository{db, lib.GetLogger()}
	defer repo.Db.Close()

	start := time.Now()
	repo.Logger.Log("METHOD", "TestUserDB", "SPOT", "method start", "time_start", start)

	//GetAll
	users, _, err := repo.GetAll("1", "0", "created_at", "desc")
	expectedInt := 1
	if expectedInt != len(users) {
		t.Errorf("\n...expected = %v\n...obtained = %v", expectedInt, len(users))
	}
	expected := interface{}(nil)
	if expected != err {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, err)
	}

	repo.Logger.Log("METHOD", "TestPlanDB", "SPOT", "method end", "time_spent", time.Since(start))
}
