package public

import (
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	goKitLog "github.com/go-kit/kit/log"
	lib "github.com/syedomair/plan-api/lib"
	"github.com/syedomair/plan-api/models"
)

type PublicEnv struct {
	Logger     goKitLog.Logger
	PublicRepo PublicRepositoryInterface
	Common     lib.CommonService
}

func (env *PublicEnv) Ping(w http.ResponseWriter, r *http.Request) {

	env.Logger.Log("METHOD", "Ping", "SPOT", "method start")
	start := time.Now()

	responseToken := map[string]string{"response": "pong"}

	env.Logger.Log("METHOD", "Ping", "SPOT", "METHOD END", "time_spent", time.Since(start))
	env.Common.SuccessResponseHelper(w, responseToken, http.StatusCreated)
}

func (env *PublicEnv) Register(w http.ResponseWriter, r *http.Request) {

	env.Logger.Log("METHOD", "Register", "SPOT", "method start")
	start := time.Now()

	var paramConf map[string]models.ParamConf
	paramConf = make(map[string]models.ParamConf)
	paramConf["first_name"] = models.ParamConf{Required: true, Type: lib.STRING_NAME, EmptyAllowed: false}
	paramConf["last_name"] = models.ParamConf{Required: true, Type: lib.STRING_NAME, EmptyAllowed: false}
	paramConf["email"] = models.ParamConf{Required: true, Type: lib.STRING_EMAIL, EmptyAllowed: false}
	paramConf["password"] = models.ParamConf{Required: true, Type: lib.STRING_PASSWORD, EmptyAllowed: false}

	_, paramMap, errCode, err := env.Common.ValidateInputParameters(r, paramConf, nil)
	if err != nil {
		env.Common.ErrorResponseHelper(w, "1"+errCode, err.Error(), http.StatusBadRequest)
		return
	}

	err = env.PublicRepo.IsEmailUnique(paramMap["email"].(string))
	if err != nil {
		env.Common.ErrorResponseHelper(w, "1001", err.Error(), http.StatusBadRequest)
		return
	}

	user, err := env.PublicRepo.CreateUser(paramMap)
	if err != nil {
		env.Common.ErrorResponseHelper(w, "1002", lib.ERROR_UNEXPECTED, http.StatusInternalServerError)
		return
	}
	env.Logger.Log("METHOD", "Register", "SPOT", "Register", "userId", user.Id)

	signedJwtToken := createUserToken(user)
	responseToken := map[string]string{"token": signedJwtToken}
	err = env.PublicRepo.CreateUserLogin(user.Id, signedJwtToken)
	if err != nil {
		env.Common.ErrorResponseHelper(w, "1003", lib.ERROR_UNEXPECTED, http.StatusInternalServerError)
		return
	}
	env.Logger.Log("METHOD", "Register", "SPOT", "METHOD END", "time_spent", time.Since(start))
	env.Common.SuccessResponseHelper(w, responseToken, http.StatusCreated)
}

func (env *PublicEnv) Login(w http.ResponseWriter, r *http.Request) {

	env.Logger.Log("METHOD", "Login", "SPOT", "method start")
	start := time.Now()

	var paramConf map[string]models.ParamConf
	paramConf = make(map[string]models.ParamConf)
	paramConf["email"] = models.ParamConf{Required: true, Type: lib.STRING_EMAIL, EmptyAllowed: false}
	paramConf["password"] = models.ParamConf{Required: true, Type: lib.STRING_PASSWORD, EmptyAllowed: false}

	_, paramMap, errCode, err := env.Common.ValidateInputParameters(r, paramConf, nil)
	if err != nil {
		env.Common.ErrorResponseHelper(w, "1"+errCode, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := env.PublicRepo.ValidateEmailPasswordFromDB(paramMap["email"].(string), paramMap["password"].(string))
	if err != nil {
		env.Common.ErrorResponseHelper(w, "1004", "Invalid email or password", http.StatusUnauthorized)
		return
	}

	signedJwtToken := createUserToken(user)
	responseToken := map[string]string{"token": signedJwtToken}

	err = env.PublicRepo.CreateUserLogin(user.Id, signedJwtToken)
	if err != nil {
		env.Common.ErrorResponseHelper(w, "1005", lib.ERROR_UNEXPECTED, http.StatusInternalServerError)
		return
	}
	env.Logger.Log("METHOD", "Login", "SPOT", "METHOD END", "time_spent", time.Since(start))
	env.Common.SuccessResponseHelper(w, responseToken, http.StatusOK)
}

func createUserToken(user *models.User) string {
	type Claims struct {
		CurrentUserId string `json:"current_user_id"`
		FirstName     string `json:"first_name"`
		LastName      string `json:"last_name"`
		Email         string `json:"email"`
		CreatedAt     int64  `json:"created_at"`
		jwt.StandardClaims
	}

	claims := Claims{
		user.Id,
		user.FirstName,
		user.LastName,
		user.Email,
		time.Now().UnixNano(),
		jwt.StandardClaims{
			Issuer: "SYEDOMAIR",
		},
	}
	signingKey := []byte(lib.SIGNING_KEY)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedJwtToken, _ := token.SignedString(signingKey)
	return signedJwtToken
}

//VIM command for serializing error code :let @a=1001 | %s/\d\d\d\d/\=''.(@a+setreg('a',@a+1))/g
