package public

import (
	"strconv"
	"testing"
	"time"

	"github.com/syedomair/plan-api/lib"
	testdata "github.com/syedomair/plan-api/test/testdata"
)

func TestPublicDB(t *testing.T) {

	db, _ := lib.CreateDBConnection()
	repo := &PublicRepository{db, lib.GetLogger()}
	defer repo.Db.Close()

	start := time.Now()
	repo.Logger.Log("METHOD", "TestPublicDB", "SPOT", "method start", "time_start", start)

	currentTime := time.Now()
	uniqueEmail := "email_" + strconv.FormatInt(currentTime.UnixNano(), 10) + "@gmail.com"
	userInput := make(map[string]interface{})
	userInput["first_name"] = testdata.ValidFirstName
	userInput["last_name"] = testdata.ValidLastName
	userInput["email"] = uniqueEmail
	userInput["password"] = testdata.ValidPassword

	user, err := repo.CreateUser(userInput)
	var expected error = nil
	if expected != err {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, err)
	}
	repo.Logger.Log("METHOD", "TestPublicDB", "userId", user.Id)

	err = repo.IsEmailUnique(uniqueEmail)
	expectedStr := "User email already exists."
	if expectedStr != err.Error() {
		t.Errorf("\n...expected = %v\n...obtained = %v", expectedStr, err)
	}

	err = repo.CreateUserLogin(user.Id, testdata.Token)
	expected = nil
	if expected != err {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, err)
	}

	if err = repo.Db.Delete(&user).Error; err != nil {
		repo.Logger.Log("METHOD", "TestPublicDB", "Error in deleting", err)
	}
	repo.Logger.Log("METHOD", "TestPublicDB", "SPOT", "method end", "time_spent", time.Since(start))
}
