package plnmsg

import (
	"testing"
	"time"

	"github.com/syedomair/plan-api/lib"
	"github.com/syedomair/plan-api/models"
	planPkg "github.com/syedomair/plan-api/plans-svc/plan"
	testdata "github.com/syedomair/plan-api/test/testdata"
)

func TestPlanMsgDB(t *testing.T) {

	db, _ := lib.CreateDBConnection()
	repoPlan := &planPkg.PlanRepository{db, lib.GetLogger()}
	defer repoPlan.Db.Close()

	//Create Plan
	plan := make(map[string]interface{})
	plan["title"] = testdata.ValidPlanTitle
	plan["status"] = testdata.ValidPlanStatus
	plan["cost"] = testdata.ValidPlanCost
	plan["validity"] = testdata.ValidPlanValidity
	planId, err := repoPlan.Create(plan)

	repo := &PlanMessageRepository{db, lib.GetLogger()}
	defer repo.Db.Close()

	start := time.Now()
	repo.Logger.Log("METHOD", "TestPlanMsgDB", "SPOT", "method start", "time_start", start)

	//Create
	planMsg := make(map[string]interface{})
	planMsg["message"] = testdata.ValidPlanMessage
	planMsg["action"] = testdata.ValidPlanMessageAction
	planMsgId, err := repo.Create(planMsg, planId)
	var expected error = nil
	if expected != err {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, err)
	}
	repo.Logger.Log("METHOD", "TestPlanMsgDB", "planMsgId", planMsgId)

	//Get
	planResponse, err := repo.Get(planMsgId)
	expected = nil
	if expected != err {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, err)
	}
	repo.Logger.Log("METHOD", "TestPlanMsgDB", "planResponse", planResponse)

	//Update
	planMsg["message"] = testdata.ValidPlanMessage + "changed"
	err = repo.Update(planMsg, planMsgId)
	expected = nil
	if expected != err {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, err)
	}

	//Get
	planResponse, err = repo.Get(planMsgId)
	expected = nil
	if expected != err {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, err)
	}
	expectedString := testdata.ValidPlanMessage + "changed"
	if expectedString != planResponse.Message {
		t.Errorf("\n...expected = %v\n...obtained = %v", expectedString, planResponse.Message)
	}
	repo.Logger.Log("METHOD", "TestPlanMsgDB", "planResponse", planResponse)

	//GetAll
	plans, _, err := repo.GetAll("1", "0", "created_at", "desc", planId)
	expectedInt := 1
	if expectedInt != len(plans) {
		t.Errorf("\n...expected = %v\n...obtained = %v", expectedInt, len(plans))
	}
	expected = nil
	if expected != err {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, err)
	}

	//cleaning up after test
	planMsgModel := models.PlanMessage{}
	planMsgModel.Id = planMsgId

	err = repo.Delete(planMsgModel)
	expected = nil
	if expected != err {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, err)
	}

	//cleaning up after test
	planModel := models.Plan{}
	planModel.Id = planId

	err = repoPlan.Delete(planModel)
	expected = nil
	if expected != err {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, err)
	}
	repo.Logger.Log("METHOD", "TestPlanMsgDB", "SPOT", "method end", "time_spent", time.Since(start))
}
