package public

import (
	"encoding/json"
	"strconv"
	"testing"
	"time"

	lib "github.com/syedomair/plan-api/lib"
	"github.com/syedomair/plan-api/models"
	"github.com/syedomair/plan-api/test/testdata"
)

func TestRegister(t *testing.T) {
	env := PublicEnv{Logger: lib.GetLogger(), PublicRepo: &mockRepo{}, Common: lib.CommonService{Logger: lib.GetLogger()}}
	method := "POST"
	url := "/register"
	currentTime := time.Now()
	uniqueEmail := "email_" + strconv.FormatInt(currentTime.UnixNano(), 10) + "@gmail.com"

	//Invalid first name
	res, req := env.Common.MockTestServer(method, url, []byte(`{"first_name":"`+testdata.InValidFirstName+`" , "last_name":"`+testdata.ValidLastName+`"  , "email":"`+uniqueEmail+`", "password":"`+testdata.ValidPassword+`" }`))
	env.Register(res, req)
	response := new(models.TestResponse)
	json.NewDecoder(res.Result().Body).Decode(response)
	expected := testdata.FAILURE
	if expected != response.Result {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, response.Result)
	}

	//Invalid last name
	res, req = env.Common.MockTestServer(method, url, []byte(`{"first_name":"`+testdata.ValidFirstName+`" , "last_name":"`+testdata.InValidLastName+`"  , "email":"`+uniqueEmail+`", "password":"`+testdata.ValidPassword+`" }`))
	env.Register(res, req)
	response = new(models.TestResponse)
	json.NewDecoder(res.Result().Body).Decode(response)
	expected = testdata.FAILURE
	if expected != response.Result {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, response.Result)
	}

	//Invalid email
	res, req = env.Common.MockTestServer(method, url, []byte(`{"first_name":"`+testdata.ValidFirstName+`" , "last_name":"`+testdata.ValidLastName+`"  , "email":"`+testdata.InValidEmailBlank+`", "password":"`+testdata.ValidPassword+`" }`))
	env.Register(res, req)
	response = new(models.TestResponse)
	json.NewDecoder(res.Result().Body).Decode(response)
	expected = testdata.FAILURE
	if expected != response.Result {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, response.Result)
	}

	//missing email
	res, req = env.Common.MockTestServer(method, url, []byte(`{"first_name":"`+testdata.ValidFirstName+`" , "last_name":"`+testdata.ValidLastName+`"  , "password":"`+testdata.ValidPassword+`" }`))
	env.Register(res, req)
	response = new(models.TestResponse)
	json.NewDecoder(res.Result().Body).Decode(response)
	expected = testdata.FAILURE
	if expected != response.Result {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, response.Result)
	}

	//Invalid password
	res, req = env.Common.MockTestServer(method, url, []byte(`{"first_name":"`+testdata.ValidFirstName+`" , "last_name":"`+testdata.ValidLastName+`"  , "email":"`+uniqueEmail+`", "password":"`+testdata.InValidPasswordBlank+`" }`))
	env.Register(res, req)
	response = new(models.TestResponse)
	json.NewDecoder(res.Result().Body).Decode(response)
	expected = testdata.FAILURE
	if expected != response.Result {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, response.Result)
	}

	//All Good
	res, req = env.Common.MockTestServer(method, url, []byte(`{"first_name":"`+testdata.ValidFirstName+`" , "last_name":"`+testdata.ValidLastName+`"  , "email":"`+uniqueEmail+`", "password":"`+testdata.ValidPassword+`" }`))
	env.Register(res, req)
	response = new(models.TestResponse)
	json.NewDecoder(res.Result().Body).Decode(response)
	expected = testdata.SUCCESS
	if expected != response.Result {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, response.Result)
	}
}

func TestLogin(t *testing.T) {
	env := PublicEnv{Logger: lib.GetLogger(), PublicRepo: &mockRepo{}, Common: lib.CommonService{Logger: lib.GetLogger()}}
	method := "POST"
	url := "/login"

	//Invalid email
	res, req := env.Common.MockTestServer(method, url, []byte(`{"email":"`+testdata.InValidEmail+`" , "password":"`+testdata.ValidPassword+`" }`))
	env.Login(res, req)
	response := new(models.TestResponse)
	json.NewDecoder(res.Result().Body).Decode(response)
	expected := testdata.FAILURE
	if expected != response.Result {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, response.Result)
	}

	//Invalid password
	res, req = env.Common.MockTestServer(method, url, []byte(`{"email":"`+testdata.ValidEmail+`" , "password":"`+testdata.InValidPasswordBlank+`" }`))
	env.Login(res, req)
	response = new(models.TestResponse)
	json.NewDecoder(res.Result().Body).Decode(response)
	expected = testdata.FAILURE
	if expected != response.Result {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, response.Result)
	}

	//All Good
	res, req = env.Common.MockTestServer(method, url, []byte(`{"email":"`+testdata.ValidEmail+`" , "password":"`+testdata.ValidPassword+`" }`))
	env.Login(res, req)
	response = new(models.TestResponse)
	json.NewDecoder(res.Result().Body).Decode(response)
	expected = testdata.SUCCESS
	if expected != response.Result {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, response.Result)
	}
}

type mockRepo struct {
}

func (mdb *mockRepo) CreateUser(inputUser map[string]interface{}) (*models.User, error) {
	user := models.User{}
	return &user, nil
}

func (mdb *mockRepo) IsEmailUnique(email string) error {
	return nil
}

func (mdb *mockRepo) CreateUserLogin(userId string, token string) error {
	return nil
}
func (mdb *mockRepo) ValidateEmailPasswordFromDB(email string, password string) (*models.User, error) {
	user := models.User{}
	return &user, nil
}
func (mdb *mockRepo) FindToken(token string) (models.UsersLogin, error) {
	userLogin := models.UsersLogin{}
	return userLogin, nil
}
