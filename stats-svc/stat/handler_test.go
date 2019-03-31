package stat

import (
	"encoding/json"
	"testing"

	lib "github.com/syedomair/plan-api/lib"
	"github.com/syedomair/plan-api/models"
	"github.com/syedomair/plan-api/test/testdata"
)

func TestGetTotalUserCountLast30Days(t *testing.T) {
	env := StatEnv{Logger: lib.GetLogger(), StatRepo: &mockRepo{}, Common: lib.CommonService{Logger: lib.GetLogger()}}
	method := "GET"
	url := "/stats/user-count-30-days"
	var blankByte []byte

	res, req := env.Common.MockTestServer(method, url, blankByte)
	env.GetTotalUserCountLast30Days(res, req)
	response := new(models.TestResponse)
	json.NewDecoder(res.Result().Body).Decode(response)
	expected := testdata.SUCCESS
	if expected != response.Result {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, response.Result)
	}
}

func TestGetTotalUserCount(t *testing.T) {
	env := StatEnv{Logger: lib.GetLogger(), StatRepo: &mockRepo{}, Common: lib.CommonService{Logger: lib.GetLogger()}}
	method := "GET"
	url := "/stats/user-count"
	var blankByte []byte

	res, req := env.Common.MockTestServer(method, url, blankByte)
	env.GetTotalUserCount(res, req)
	response := new(models.TestResponse)
	json.NewDecoder(res.Result().Body).Decode(response)
	expected := testdata.SUCCESS
	if expected != response.Result {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, response.Result)
	}
}
func TestGetUserRegData(t *testing.T) {
	env := StatEnv{Logger: lib.GetLogger(), StatRepo: &mockRepo{}, Common: lib.CommonService{Logger: lib.GetLogger()}}
	method := "GET"
	url := "/stats/user-reg-data"
	var blankByte []byte

	res, req := env.Common.MockTestServer(method, url, blankByte)
	env.GetUserRegData(res, req)
	response := new(models.TestResponse)
	json.NewDecoder(res.Result().Body).Decode(response)
	expected := testdata.SUCCESS
	if expected != response.Result {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, response.Result)
	}
}
func TestGetPlanData(t *testing.T) {
	env := StatEnv{Logger: lib.GetLogger(), StatRepo: &mockRepo{}, Common: lib.CommonService{Logger: lib.GetLogger()}}
	method := "GET"
	url := "/stats/plan-data"
	var blankByte []byte

	res, req := env.Common.MockTestServer(method, url, blankByte)
	env.GetPlanData(res, req)
	response := new(models.TestResponse)
	json.NewDecoder(res.Result().Body).Decode(response)
	expected := testdata.SUCCESS
	if expected != response.Result {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, response.Result)
	}
}

type mockRepo struct {
}

func (mdb *mockRepo) GetTotalUserCount() (string, error) {
	return "1", nil
}
func (mdb *mockRepo) GetTotalUserCountLast30Days() (string, error) {
	return "1", nil
}
func (mdb *mockRepo) GetUserRegData() ([]*models.StatUserRegPerMonth, error) {
	var statUserRegPerMonth []*models.StatUserRegPerMonth
	return statUserRegPerMonth, nil
}
func (mdb *mockRepo) GetPlanData() (string, error) {
	return "1", nil
}
