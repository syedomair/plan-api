package plan

import (
	"testing"
	"time"

	"github.com/syedomair/plan-api/lib"
	"github.com/syedomair/plan-api/models"
	testdata "github.com/syedomair/plan-api/test/testdata"
)

func TestPlanDB(t *testing.T) {

	db, _ := lib.CreateDBConnection()
	repo := &PlanRepository{db, lib.GetLogger()}
	defer repo.Db.Close()

	start := time.Now()
	repo.Logger.Log("METHOD", "TestPlanDB", "SPOT", "method start", "time_start", start)

	//Create
	plan := make(map[string]interface{})
	plan["title"] = testdata.ValidPlanTitle
	plan["status"] = testdata.ValidPlanStatus
	plan["cost"] = testdata.ValidPlanCost
	plan["validity"] = testdata.ValidPlanValidity
	planId, err := repo.Create(plan)
	var expected error = nil
	if expected != err {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, err)
	}
	repo.Logger.Log("METHOD", "TestPlanDB", "planId", planId)
	//Get
	planResponse, err := repo.Get(planId)
	expected = nil
	if expected != err {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, err)
	}
	repo.Logger.Log("METHOD", "TestPlanDB", "planResponse", planResponse)

	//Update
	plan["title"] = testdata.ValidPlanTitle + "changed"
	err = repo.Update(plan, planId)
	expected = nil
	if expected != err {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, err)
	}

	//Get
	planResponse, err = repo.Get(planId)
	expected = nil
	if expected != err {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, err)
	}
	expectedString := testdata.ValidPlanTitle + "changed"
	if expectedString != planResponse.Title {
		t.Errorf("\n...expected = %v\n...obtained = %v", expectedString, planResponse.Title)
	}
	repo.Logger.Log("METHOD", "TestPlanDB", "planResponse", planResponse)

	//GetAll
	plans, _, err := repo.GetAll("1", "0", "created_at", "desc")
	expectedInt := 1
	if expectedInt != len(plans) {
		t.Errorf("\n...expected = %v\n...obtained = %v", expectedInt, len(plans))
	}
	expected = nil
	if expected != err {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, err)
	}

	//cleaning up after test
	planModel := models.Plan{}
	planModel.Id = planId

	err = repo.Delete(planModel)
	expected = nil
	if expected != err {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, err)
	}
	repo.Logger.Log("METHOD", "TestPlanDB", "SPOT", "method end", "time_spent", time.Since(start))
}
