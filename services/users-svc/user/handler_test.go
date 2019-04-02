package user

import (
	"encoding/json"
	"testing"

	lib "github.com/syedomair/plan-api/lib"
	"github.com/syedomair/plan-api/models"
	"github.com/syedomair/plan-api/test/testdata"
)

func TestGetAll(t *testing.T) {
	env := UserEnv{Logger: lib.GetLogger(), UserRepo: &mockRepo{}, Common: lib.CommonService{Logger: lib.GetLogger()}}
	method := "GET"
	url := "/users"
	var blankByte []byte

	//All Good
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

func (mdb *mockRepo) GetAll(limit string, offset string, orderby string, sort string) ([]*models.UserReduced, string, error) {
	var users []*models.UserReduced
	return users, "0", nil
}
