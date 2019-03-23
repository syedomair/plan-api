package user

import (
	"net/http"
	"time"

	goKitLog "github.com/go-kit/kit/log"
	lib "github.com/syedomair/plan-api/lib"
)

type UserEnv struct {
	Logger   goKitLog.Logger
	UserRepo UserRepositoryInterface
	Common   lib.CommonService
}

func (env *UserEnv) GetAll(w http.ResponseWriter, r *http.Request) {

	start := time.Now()
	env.Logger.Log("METHOD", "GetAll", "SPOT", "method start", "time_start", start)

	_, err := env.Common.GetUserClientFromToken(r)
	if err != nil {
		env.Common.ErrorResponseHelper(w, "2001", err.Error(), http.StatusBadRequest)
		return
	}

	limit, offset, orderby, sort, err := env.Common.ValidateQueryString(r, "100", "0", "created_at", "asc")
	if err != nil {
		env.Common.ErrorResponseHelper(w, "2002", err.Error(), http.StatusBadRequest)
		return
	}
	users, count, err := env.UserRepo.GetAll(limit, offset, orderby, sort)
	if err != nil {
		env.Common.ErrorResponseHelper(w, "2003", lib.ERROR_UNEXPECTED, http.StatusInternalServerError)
		return
	}
	env.Logger.Log("METHOD", "GetAll", "SPOT", "method end", "time_spent", time.Since(start))
	env.Common.SuccessResponseList(w, users, offset, limit, count)
}
