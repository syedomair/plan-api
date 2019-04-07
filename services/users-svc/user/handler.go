package user

import (
	"net/http"
	"time"

	goKitLog "github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
	lib "github.com/syedomair/plan-api/lib"
)

type UserEnv struct {
	Logger   goKitLog.Logger
	UserRepo UserRepositoryInterface
	Common   lib.CommonService
}

func (env *UserEnv) GetAllBatch(w http.ResponseWriter, r *http.Request) {

	start := time.Now()
	env.Logger.Log("METHOD", "GetAllBatch", "SPOT", "method start", "time_start", start)

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
	batchTaskId, err := env.UserRepo.GetAllBatch(limit, offset, orderby, sort)
	if err != nil {
		env.Common.ErrorResponseHelper(w, "2003", lib.ERROR_UNEXPECTED, http.StatusInternalServerError)
		return
	}

	responseBatchTaskId := map[string]string{"batch_task_id": batchTaskId}

	env.Logger.Log("METHOD", "GetAllBatch", "SPOT", "METHOD END", "time_spent", time.Since(start))
	env.Common.SuccessResponseHelper(w, responseBatchTaskId, http.StatusCreated)
}

func (env *UserEnv) GetBatchTask(w http.ResponseWriter, r *http.Request) {

	start := time.Now()
	env.Logger.Log("METHOD", "Get", "SPOT", "method start", "time_start", start)
	_, err := env.Common.GetUserClientFromToken(r)
	if err != nil {
		env.Common.ErrorResponseHelper(w, "3006", err.Error(), http.StatusBadRequest)
		return
	}
	pathParams := mux.Vars(r)
	batchTaskId := pathParams["batch_task_id"]

	if err := env.Common.ValidateId(batchTaskId, "batch_task"); err != nil {
		env.Common.ErrorResponseHelper(w, "3007", err.Error(), http.StatusBadRequest)
		return
	}

	batchTask, err := env.UserRepo.GetBatchTask(batchTaskId)
	if err != nil {
		env.Common.ErrorResponseHelper(w, "3008", lib.ERROR_UNEXPECTED, http.StatusInternalServerError)
		return
	}
	env.Logger.Log("METHOD", "Get", "SPOT", "method end", "time_spent", time.Since(start))
	env.Common.SuccessResponseHelper(w, batchTask, http.StatusOK)
}

func (env *UserEnv) GetAll(w http.ResponseWriter, r *http.Request) {

	start := time.Now()
	env.Logger.Log("METHOD", "GetAll", "SPOT", "method start", "time_start", start)

	_, err := env.Common.GetUserClientFromToken(r)
	if err != nil {
		env.Common.ErrorResponseHelper(w, "2004", err.Error(), http.StatusBadRequest)
		return
	}

	limit, offset, orderby, sort, err := env.Common.ValidateQueryString(r, "100", "0", "created_at", "asc")
	if err != nil {
		env.Common.ErrorResponseHelper(w, "2005", err.Error(), http.StatusBadRequest)
		return
	}
	users, count, err := env.UserRepo.GetAll(limit, offset, orderby, sort)
	if err != nil {
		env.Common.ErrorResponseHelper(w, "2006", lib.ERROR_UNEXPECTED, http.StatusInternalServerError)
		return
	}
	env.Logger.Log("METHOD", "GetAll", "SPOT", "method end", "time_spent", time.Since(start))
	env.Common.SuccessResponseList(w, users, offset, limit, count)
}

//VIM command to serializing error code :let @a=2007 | %s/\d\d\d\d/\=''.(@a+setreg('a',@a+1))/g
