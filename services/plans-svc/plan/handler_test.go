package plan

import (
	"encoding/json"
	"testing"

	lib "github.com/syedomair/plan-api/lib"
	"github.com/syedomair/plan-api/models"
	"github.com/syedomair/plan-api/test/testdata"
)

func TestCreatePlan(t *testing.T) {
	env := PlanEnv{Logger: lib.GetLogger(), PlanRepo: &mockRepo{}, Common: lib.CommonService{Logger: lib.GetLogger()}}
	method := "POST"
	url := "/plans"

	//Invalid Title
	res, req := env.Common.MockTestServer(method, url, []byte(`{"title":"`+testdata.InValidLenPlanTitle+`" , "status":"`+testdata.ValidPlanStatus+`"  , "cost":"`+testdata.ValidPlanCost+`", "validity":"`+testdata.ValidPlanValidity+`" }`))
	env.CreatePlan(res, req)
	response := new(models.TestResponse)
	json.NewDecoder(res.Result().Body).Decode(response)
	expected := testdata.FAILURE
	if expected != response.Result {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, response.Result)
	}

	//Invalid invalid title blank
	res, req = env.Common.MockTestServer(method, url, []byte(`{"title":"`+testdata.InValidPlanTitleBlank+`" , "status":"`+testdata.ValidPlanStatus+`"  , "cost":"`+testdata.ValidPlanCost+`", "validity":"`+testdata.ValidPlanValidity+`" }`))
	env.CreatePlan(res, req)
	response = new(models.TestResponse)
	json.NewDecoder(res.Result().Body).Decode(response)
	expected = testdata.FAILURE
	if expected != response.Result {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, response.Result)
	}

	//Invalid invalid cost
	res, req = env.Common.MockTestServer(method, url, []byte(`{"title":"`+testdata.InValidPlanTitleBlank+`" , "status":"`+testdata.ValidPlanStatus+`"  , "cost":"`+testdata.InValidPlanCost+`", "validity":"`+testdata.ValidPlanValidity+`" }`))
	env.CreatePlan(res, req)
	response = new(models.TestResponse)
	json.NewDecoder(res.Result().Body).Decode(response)
	expected = testdata.FAILURE
	if expected != response.Result {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, response.Result)
	}

	//Invalid invalid cost blank
	res, req = env.Common.MockTestServer(method, url, []byte(`{"title":"`+testdata.InValidPlanTitleBlank+`" , "status":"`+testdata.ValidPlanStatus+`"  , "cost":"`+testdata.InValidPlanCostBlank+`", "validity":"`+testdata.ValidPlanValidity+`" }`))
	env.CreatePlan(res, req)
	response = new(models.TestResponse)
	json.NewDecoder(res.Result().Body).Decode(response)
	expected = testdata.FAILURE
	if expected != response.Result {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, response.Result)
	}

	//Invalid invalid validity
	res, req = env.Common.MockTestServer(method, url, []byte(`{"title":"`+testdata.InValidPlanTitleBlank+`" , "status":"`+testdata.ValidPlanStatus+`"  , "cost":"`+testdata.ValidPlanCost+`", "validity":"`+testdata.InValidPlanValidity+`" }`))
	env.CreatePlan(res, req)
	response = new(models.TestResponse)
	json.NewDecoder(res.Result().Body).Decode(response)
	expected = testdata.FAILURE
	if expected != response.Result {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, response.Result)
	}

	//Invalid invalid validity blank
	res, req = env.Common.MockTestServer(method, url, []byte(`{"title":"`+testdata.InValidPlanTitleBlank+`" , "status":"`+testdata.ValidPlanStatus+`"  , "cost":"`+testdata.ValidPlanCost+`", "validity":"`+testdata.InValidPlanValidityBlank+`" }`))
	env.CreatePlan(res, req)
	response = new(models.TestResponse)
	json.NewDecoder(res.Result().Body).Decode(response)
	expected = testdata.FAILURE
	if expected != response.Result {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, response.Result)
	}

	//Invalid invalid status
	res, req = env.Common.MockTestServer(method, url, []byte(`{"title":"`+testdata.InValidPlanTitleBlank+`" , "status":"`+testdata.InValidPlanStatus+`"  , "cost":"`+testdata.ValidPlanCost+`", "validity":"`+testdata.ValidPlanValidity+`" }`))
	env.CreatePlan(res, req)
	response = new(models.TestResponse)
	json.NewDecoder(res.Result().Body).Decode(response)
	expected = testdata.FAILURE
	if expected != response.Result {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, response.Result)
	}

	//Invalid invalid status blank
	res, req = env.Common.MockTestServer(method, url, []byte(`{"title":"`+testdata.InValidPlanTitleBlank+`" , "status":"`+testdata.InValidPlanStatusBlank+`"  , "cost":"`+testdata.ValidPlanCost+`", "validity":"`+testdata.ValidPlanValidity+`" }`))
	env.CreatePlan(res, req)
	response = new(models.TestResponse)
	json.NewDecoder(res.Result().Body).Decode(response)
	expected = testdata.FAILURE
	if expected != response.Result {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, response.Result)
	}
	// ALL GOOD
	res, req = env.Common.MockTestServer(method, url, []byte(`{"title":"`+testdata.ValidPlanTitle+`" , "status":"`+testdata.ValidPlanStatus+`"  , "cost":"`+testdata.ValidPlanCost+`", "validity":"`+testdata.ValidPlanValidity+`" }`))
	env.CreatePlan(res, req)
	response = new(models.TestResponse)
	json.NewDecoder(res.Result().Body).Decode(response)
	expected = testdata.SUCCESS
	if expected != response.Result {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, response.Result)
	}
}

func TestUpdate(t *testing.T) {
	env := PlanEnv{Logger: lib.GetLogger(), PlanRepo: &mockRepo{}, Common: lib.CommonService{Logger: lib.GetLogger()}}
	method := "PATCH"
	url := "/plans/e9e0c672-64c6-488b-868f-90d695abd674"

	//Invalid Title
	res, req := env.Common.MockTestServer(method, url, []byte(`{"title":"`+testdata.InValidLenPlanTitle+`" , "status":"`+testdata.ValidPlanStatus+`"  , "cost":"`+testdata.ValidPlanCost+`", "validity":"`+testdata.ValidPlanValidity+`" }`))
	env.Update(res, req)
	response := new(models.TestResponse)
	json.NewDecoder(res.Result().Body).Decode(response)
	expected := testdata.FAILURE
	if expected != response.Result {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, response.Result)
	}

	//Invalid invalid title blank
	res, req = env.Common.MockTestServer(method, url, []byte(`{"title":"`+testdata.InValidPlanTitleBlank+`" , "status":"`+testdata.ValidPlanStatus+`"  , "cost":"`+testdata.ValidPlanCost+`", "validity":"`+testdata.ValidPlanValidity+`" }`))
	env.Update(res, req)
	response = new(models.TestResponse)
	json.NewDecoder(res.Result().Body).Decode(response)
	expected = testdata.FAILURE
	if expected != response.Result {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, response.Result)
	}

	//Invalid invalid cost
	res, req = env.Common.MockTestServer(method, url, []byte(`{"title":"`+testdata.InValidPlanTitleBlank+`" , "status":"`+testdata.ValidPlanStatus+`"  , "cost":"`+testdata.InValidPlanCost+`", "validity":"`+testdata.ValidPlanValidity+`" }`))
	env.Update(res, req)
	response = new(models.TestResponse)
	json.NewDecoder(res.Result().Body).Decode(response)
	expected = testdata.FAILURE
	if expected != response.Result {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, response.Result)
	}

	//Invalid invalid cost blank
	res, req = env.Common.MockTestServer(method, url, []byte(`{"title":"`+testdata.InValidPlanTitleBlank+`" , "status":"`+testdata.ValidPlanStatus+`"  , "cost":"`+testdata.InValidPlanCostBlank+`", "validity":"`+testdata.ValidPlanValidity+`" }`))
	env.Update(res, req)
	response = new(models.TestResponse)
	json.NewDecoder(res.Result().Body).Decode(response)
	expected = testdata.FAILURE
	if expected != response.Result {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, response.Result)
	}

	//Invalid invalid validity
	res, req = env.Common.MockTestServer(method, url, []byte(`{"title":"`+testdata.InValidPlanTitleBlank+`" , "status":"`+testdata.ValidPlanStatus+`"  , "cost":"`+testdata.ValidPlanCost+`", "validity":"`+testdata.InValidPlanValidity+`" }`))
	env.Update(res, req)
	response = new(models.TestResponse)
	json.NewDecoder(res.Result().Body).Decode(response)
	expected = testdata.FAILURE
	if expected != response.Result {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, response.Result)
	}

	//Invalid invalid validity blank
	res, req = env.Common.MockTestServer(method, url, []byte(`{"title":"`+testdata.InValidPlanTitleBlank+`" , "status":"`+testdata.ValidPlanStatus+`"  , "cost":"`+testdata.ValidPlanCost+`", "validity":"`+testdata.InValidPlanValidityBlank+`" }`))
	env.Update(res, req)
	response = new(models.TestResponse)
	json.NewDecoder(res.Result().Body).Decode(response)
	expected = testdata.FAILURE
	if expected != response.Result {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, response.Result)
	}

	//Invalid invalid status
	res, req = env.Common.MockTestServer(method, url, []byte(`{"title":"`+testdata.InValidPlanTitleBlank+`" , "status":"`+testdata.InValidPlanStatus+`"  , "cost":"`+testdata.ValidPlanCost+`", "validity":"`+testdata.ValidPlanValidity+`" }`))
	env.Update(res, req)
	response = new(models.TestResponse)
	json.NewDecoder(res.Result().Body).Decode(response)
	expected = testdata.FAILURE
	if expected != response.Result {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, response.Result)
	}

	//Invalid invalid status blank
	res, req = env.Common.MockTestServer(method, url, []byte(`{"title":"`+testdata.InValidPlanTitleBlank+`" , "status":"`+testdata.InValidPlanStatusBlank+`"  , "cost":"`+testdata.ValidPlanCost+`", "validity":"`+testdata.ValidPlanValidity+`" }`))
	env.Update(res, req)
	response = new(models.TestResponse)
	json.NewDecoder(res.Result().Body).Decode(response)
	expected = testdata.FAILURE
	if expected != response.Result {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, response.Result)
	}
	// ALL GOOD
	res, req = env.Common.MockTestServer(method, url, []byte(`{"title":"`+testdata.ValidPlanTitle+`" , "status":"`+testdata.ValidPlanStatus+`"  , "cost":"`+testdata.ValidPlanCost+`", "validity":"`+testdata.ValidPlanValidity+`" }`))
	env.Update(res, req)
	response = new(models.TestResponse)
	json.NewDecoder(res.Result().Body).Decode(response)
	expected = testdata.SUCCESS
	if expected != response.Result {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, response.Result)
	}
}

func TestGet(t *testing.T) {
	env := PlanEnv{Logger: lib.GetLogger(), PlanRepo: &mockRepo{}, Common: lib.CommonService{Logger: lib.GetLogger()}}
	method := "GET"
	url := "/plans/invalidId"
	var blankByte []byte

	//Invalid ID
	res, req := env.Common.MockTestServer(method, url, blankByte)
	env.Get(res, req)
	response := new(models.TestResponse)
	json.NewDecoder(res.Result().Body).Decode(response)
	expected := testdata.SUCCESS
	if expected != response.Result {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, response.Result)
	}
	//ALL GOOD
	url = "/plans/e9e0c672-64c6-488b-868f-90d695abd674"
	res, req = env.Common.MockTestServer(method, url, blankByte)
	env.Get(res, req)
	response = new(models.TestResponse)
	json.NewDecoder(res.Result().Body).Decode(response)
	expected = testdata.SUCCESS
	if expected != response.Result {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, response.Result)
	}
}

func TestDelete(t *testing.T) {
	env := PlanEnv{Logger: lib.GetLogger(), PlanRepo: &mockRepo{}, Common: lib.CommonService{Logger: lib.GetLogger()}}
	method := "DELETE"
	url := "/plans/invalidId"
	var blankByte []byte

	//Invalid ID
	res, req := env.Common.MockTestServer(method, url, blankByte)
	env.Delete(res, req)
	response := new(models.TestResponse)
	json.NewDecoder(res.Result().Body).Decode(response)
	expected := testdata.SUCCESS
	if expected != response.Result {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, response.Result)
	}
	//ALL GOOD
	url = "/plans/e9e0c672-64c6-488b-868f-90d695abd674"
	res, req = env.Common.MockTestServer(method, url, blankByte)
	env.Delete(res, req)
	response = new(models.TestResponse)
	json.NewDecoder(res.Result().Body).Decode(response)
	expected = testdata.SUCCESS
	if expected != response.Result {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, response.Result)
	}
}

func TestGetAll(t *testing.T) {
	env := PlanEnv{Logger: lib.GetLogger(), PlanRepo: &mockRepo{}, Common: lib.CommonService{Logger: lib.GetLogger()}}
	method := "GET"
	url := "/plans"
	var blankByte []byte

	//ALL GOOD
	res, req := env.Common.MockTestServer(method, url, blankByte)
	env.GetAll(res, req)
	response := new(models.TestResponse)
	json.NewDecoder(res.Result().Body).Decode(response)
	expected := testdata.SUCCESS
	if expected != response.Result {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, response.Result)
	}
}

type mockRepo struct {
}

func (mdb *mockRepo) Create(inputPlan map[string]interface{}) (string, error) {
	return "e9e0c672-64c6-488b-868f-90d695abd674", nil
}
func (mdb *mockRepo) GetAll(limit string, offset string, orderby string, sort string) ([]*models.Plan, string, error) {
	var plans []*models.Plan
	return plans, "0", nil
}
func (mdb *mockRepo) Get(planId string) (*models.Plan, error) {
	plan := models.Plan{}

	return &plan, nil
}
func (mdb *mockRepo) Update(inputPlan map[string]interface{}, planId string) error {
	return nil
}
func (mdb *mockRepo) Delete(plan models.Plan) error {
	return nil
}
