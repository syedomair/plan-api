package stat

import (
	"net/http"
	"time"

	goKitLog "github.com/go-kit/kit/log"
	lib "github.com/syedomair/plan-api/lib"
)

type StatEnv struct {
	Logger   goKitLog.Logger
	StatRepo StatRepositoryInterface
	Common   lib.CommonService
}

func (env *StatEnv) GetTotalUserCount(w http.ResponseWriter, r *http.Request) {

	start := time.Now()
	env.Logger.Log("METHOD", "GetTotalUserCount", "SPOT", "method start", "time_start", start)

	_, err := env.Common.GetUserClientFromToken(r)
	if err != nil {
		env.Common.ErrorResponseHelper(w, "5001", err.Error(), http.StatusBadRequest)
		return
	}

	userCount, err := env.StatRepo.GetTotalUserCount()
	if err != nil {
		env.Common.ErrorResponseHelper(w, "5002", lib.ERROR_UNEXPECTED, http.StatusInternalServerError)
		return
	}
	responseUserCount := map[string]string{"user_total_count": userCount}
	env.Logger.Log("METHOD", "GetTotalUserCount", "SPOT", "method end", "time_spent", time.Since(start))
	env.Common.SuccessResponseHelper(w, responseUserCount, http.StatusOK)

}

func (env *StatEnv) GetUserRegData(w http.ResponseWriter, r *http.Request) {

	start := time.Now()
	env.Logger.Log("METHOD", "GetUserRegData", "SPOT", "method start", "time_start", start)

	_, err := env.Common.GetUserClientFromToken(r)
	if err != nil {
		env.Common.ErrorResponseHelper(w, "5003", err.Error(), http.StatusBadRequest)
		return
	}

	userRegData, err := env.StatRepo.GetUserRegData()
	if err != nil {
		env.Common.ErrorResponseHelper(w, "5004", lib.ERROR_UNEXPECTED, http.StatusInternalServerError)
		return
	}
	env.Logger.Log("METHOD", "GetUserRegData", "SPOT", "method end", "time_spent", time.Since(start))
	env.Common.SuccessResponseHelper(w, userRegData, http.StatusOK)

}

func (env *StatEnv) GetPlanData(w http.ResponseWriter, r *http.Request) {

	start := time.Now()
	env.Logger.Log("METHOD", "GetPlanData", "SPOT", "method start", "time_start", start)

	_, err := env.Common.GetUserClientFromToken(r)
	if err != nil {
		env.Common.ErrorResponseHelper(w, "5005", err.Error(), http.StatusBadRequest)
		return
	}

	userCount, err := env.StatRepo.GetPlanData()
	if err != nil {
		env.Common.ErrorResponseHelper(w, "5006", lib.ERROR_UNEXPECTED, http.StatusInternalServerError)
		return
	}
	responseUserCount := map[string]string{"plan_total_count": userCount}
	env.Logger.Log("METHOD", "GetPlanData", "SPOT", "method end", "time_spent", time.Since(start))
	env.Common.SuccessResponseHelper(w, responseUserCount, http.StatusOK)

}
