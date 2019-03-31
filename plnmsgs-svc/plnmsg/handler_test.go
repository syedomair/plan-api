package plnmsg

import (
	"encoding/json"
	"testing"

	lib "github.com/syedomair/plan-api/lib"
	"github.com/syedomair/plan-api/models"
	"github.com/syedomair/plan-api/test/testdata"
)

func TestCreatePlanMsg(t *testing.T) {
	env := PlanMessageEnv{Logger: lib.GetLogger(), PlanMessageRepo: &mockRepo{}, Common: lib.CommonService{Logger: lib.GetLogger()}}
	method := "POST"
	url := "/plan-messages"

	//Invalid Message Blank
	res, req := env.Common.MockTestServer(method, url, []byte(`{"message":"`+testdata.InValidPlanMessageBlank+`" , "action":"`+testdata.ValidPlanMessageAction+`"  }`))
	env.Create(res, req)
	response := new(models.TestResponse)
	json.NewDecoder(res.Result().Body).Decode(response)
	expected := testdata.FAILURE
	if expected != response.Result {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, response.Result)
	}

	//Invalid missing Message
	res, req = env.Common.MockTestServer(method, url, []byte(`{"action":"`+testdata.ValidPlanMessageAction+`"  }`))
	env.Create(res, req)
	response = new(models.TestResponse)
	json.NewDecoder(res.Result().Body).Decode(response)
	expected = testdata.FAILURE
	if expected != response.Result {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, response.Result)
	}

	//Invalid invalid Action Blank
	res, req = env.Common.MockTestServer(method, url, []byte(`{"message":"`+testdata.ValidPlanMessage+`" , "action":"`+testdata.InValidPlanMessageActionBlank+`"  }`))
	env.Create(res, req)
	response = new(models.TestResponse)
	json.NewDecoder(res.Result().Body).Decode(response)
	expected = testdata.FAILURE
	if expected != response.Result {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, response.Result)
	}

	//Invalid invalid Action
	res, req = env.Common.MockTestServer(method, url, []byte(`{"message":"`+testdata.ValidPlanMessage+`" , "action":"`+testdata.ValidPlanMessageAction+"test"+`"  }`))
	env.Create(res, req)
	response = new(models.TestResponse)
	json.NewDecoder(res.Result().Body).Decode(response)
	expected = testdata.FAILURE
	if expected != response.Result {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, response.Result)
	}

	//Invalid missing Action
	res, req = env.Common.MockTestServer(method, url, []byte(`{"message":"`+testdata.ValidPlanMessage+`"}`))
	env.Create(res, req)
	response = new(models.TestResponse)
	json.NewDecoder(res.Result().Body).Decode(response)
	expected = testdata.FAILURE
	if expected != response.Result {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, response.Result)
	}

	//All Good
	res, req = env.Common.MockTestServer(method, url, []byte(`{"message":"`+testdata.ValidPlanMessage+`" , "action":"`+testdata.ValidPlanMessageAction+`"  }`))
	env.Create(res, req)
	response = new(models.TestResponse)
	json.NewDecoder(res.Result().Body).Decode(response)
	expected = testdata.SUCCESS
	if expected != response.Result {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, response.Result)
	}
}

func TestUpdatePlanMsg(t *testing.T) {
	env := PlanMessageEnv{Logger: lib.GetLogger(), PlanMessageRepo: &mockRepo{}, Common: lib.CommonService{Logger: lib.GetLogger()}}
	method := "PATCH"
	url := "/plan-messages/e9e0c672-64c6-488b-868f-90d695abd674"

	//Invalid Message Blank
	res, req := env.Common.MockTestServer(method, url, []byte(`{"message":"`+testdata.InValidPlanMessageBlank+`" , "action":"`+testdata.ValidPlanMessageAction+`"  }`))
	env.Update(res, req)
	response := new(models.TestResponse)
	json.NewDecoder(res.Result().Body).Decode(response)
	expected := testdata.FAILURE
	if expected != response.Result {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, response.Result)
	}

	//Invalid missing Message
	res, req = env.Common.MockTestServer(method, url, []byte(`{"action":"`+testdata.ValidPlanMessageAction+`"  }`))
	env.Update(res, req)
	response = new(models.TestResponse)
	json.NewDecoder(res.Result().Body).Decode(response)
	expected = testdata.FAILURE
	if expected != response.Result {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, response.Result)
	}

	//Invalid invalid Action Blank
	res, req = env.Common.MockTestServer(method, url, []byte(`{"message":"`+testdata.ValidPlanMessage+`" , "action":"`+testdata.InValidPlanMessageActionBlank+`"  }`))
	env.Update(res, req)
	response = new(models.TestResponse)
	json.NewDecoder(res.Result().Body).Decode(response)
	expected = testdata.FAILURE
	if expected != response.Result {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, response.Result)
	}

	//Invalid invalid Action
	res, req = env.Common.MockTestServer(method, url, []byte(`{"message":"`+testdata.ValidPlanMessage+`" , "action":"`+testdata.ValidPlanMessageAction+"test"+`"  }`))
	env.Update(res, req)
	response = new(models.TestResponse)
	json.NewDecoder(res.Result().Body).Decode(response)
	expected = testdata.FAILURE
	if expected != response.Result {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, response.Result)
	}

	//Invalid missing Action
	res, req = env.Common.MockTestServer(method, url, []byte(`{"message":"`+testdata.ValidPlanMessage+`"}`))
	env.Update(res, req)
	response = new(models.TestResponse)
	json.NewDecoder(res.Result().Body).Decode(response)
	expected = testdata.FAILURE
	if expected != response.Result {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, response.Result)
	}

	//All Good
	res, req = env.Common.MockTestServer(method, url, []byte(`{"message":"`+testdata.ValidPlanMessage+`" , "action":"`+testdata.ValidPlanMessageAction+`"  }`))
	env.Update(res, req)
	response = new(models.TestResponse)
	json.NewDecoder(res.Result().Body).Decode(response)
	expected = testdata.SUCCESS
	if expected != response.Result {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, response.Result)
	}
}

func TestGet(t *testing.T) {
	env := PlanMessageEnv{Logger: lib.GetLogger(), PlanMessageRepo: &mockRepo{}, Common: lib.CommonService{Logger: lib.GetLogger()}}
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
	env := PlanMessageEnv{Logger: lib.GetLogger(), PlanMessageRepo: &mockRepo{}, Common: lib.CommonService{Logger: lib.GetLogger()}}
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
	env := PlanMessageEnv{Logger: lib.GetLogger(), PlanMessageRepo: &mockRepo{}, Common: lib.CommonService{Logger: lib.GetLogger()}}
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

func (mdb *mockRepo) Create(inputPlanMsg map[string]interface{}, planId string) (string, error) {
	return "e9e0c672-64c6-488b-868f-90d695abd674", nil
}
func (mdb *mockRepo) GetAll(limit string, offset string, orderby string, sort string, planId string) ([]*models.PlanMessage, string, error) {
	var planMesages []*models.PlanMessage
	return planMesages, "0", nil
}
func (mdb *mockRepo) Get(planMsgId string) (*models.PlanMessage, error) {
	planMessage := models.PlanMessage{}

	return &planMessage, nil
}
func (mdb *mockRepo) Update(inputPlanMsg map[string]interface{}, planMsgId string) error {
	return nil
}
func (mdb *mockRepo) Delete(plan models.PlanMessage) error {
	return nil
}
