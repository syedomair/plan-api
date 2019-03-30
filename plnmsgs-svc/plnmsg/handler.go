package plnmsg

import (
	"net/http"
	"time"

	goKitLog "github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
	lib "github.com/syedomair/plan-api/lib"
	"github.com/syedomair/plan-api/models"
)

type PlanMessageEnv struct {
	Logger          goKitLog.Logger
	PlanMessageRepo PlanMessageRepositoryInterface
	Common          lib.CommonService
}

func (env *PlanMessageEnv) Create(w http.ResponseWriter, r *http.Request) {

	env.Logger.Log("METHOD", "Create", "SPOT", "method start")
	start := time.Now()

	_, err := env.Common.GetUserClientFromToken(r)
	if err != nil {
		env.Common.ErrorResponseHelper(w, "4001", err.Error(), http.StatusBadRequest)
		return
	}

	var pathParamConf map[string]string
	pathParamConf = make(map[string]string)
	pathParamConf["plan_id"] = ""

	var paramConf map[string]models.ParamConf
	paramConf = make(map[string]models.ParamConf)
	paramConf["message"] = models.ParamConf{Required: true, Type: lib.STRING_LARGE, EmptyAllowed: false}
	paramConf["action"] = models.ParamConf{Required: true, Type: lib.STRING_ACTION_NAME, EmptyAllowed: false}

	pathParamValue, paramMap, errCode, err := env.Common.ValidateInputParameters(r, paramConf, pathParamConf)
	if err != nil {
		env.Common.ErrorResponseHelper(w, "6"+errCode, err.Error(), http.StatusBadRequest)
		return
	}

	planId := pathParamValue["plan_id"]

	planMessageId, err := env.PlanMessageRepo.Create(paramMap, planId)
	if err != nil {
		env.Common.ErrorResponseHelper(w, "4002", lib.ERROR_UNEXPECTED, http.StatusInternalServerError)
	}

	responsePlanMessageId := map[string]string{"plan_message_id": planMessageId}

	env.Logger.Log("METHOD", "Create", "SPOT", "METHOD END", "time_spent", time.Since(start))
	env.Common.SuccessResponseHelper(w, responsePlanMessageId, http.StatusCreated)

}

func (env *PlanMessageEnv) GetAll(w http.ResponseWriter, r *http.Request) {

	start := time.Now()
	env.Logger.Log("METHOD", "GetAll", "SPOT", "method start", "time_start", start)

	_, err := env.Common.GetUserClientFromToken(r)
	if err != nil {
		env.Common.ErrorResponseHelper(w, "4003", err.Error(), http.StatusBadRequest)
		return
	}
	pathParams := mux.Vars(r)
	planId := pathParams["plan_id"]
	if err := env.Common.ValidateId(planId, "plan_id"); err != nil {
		env.Common.ErrorResponseHelper(w, "4004", err.Error(), http.StatusBadRequest)
		return
	}

	limit, offset, orderby, sort, err := env.Common.ValidateQueryString(r, "100", "0", "created_at", "asc")
	if err != nil {
		env.Common.ErrorResponseHelper(w, "4005", err.Error(), http.StatusBadRequest)
		return
	}
	planMessages, count, err := env.PlanMessageRepo.GetAll(limit, offset, orderby, sort, planId)
	if err != nil {
		env.Common.ErrorResponseHelper(w, "4006", lib.ERROR_UNEXPECTED, http.StatusInternalServerError)
	}
	env.Logger.Log("METHOD", "GetAll", "SPOT", "method end", "time_spent", time.Since(start))
	env.Common.SuccessResponseList(w, planMessages, offset, limit, count)
}
func (env *PlanMessageEnv) Get(w http.ResponseWriter, r *http.Request) {

	start := time.Now()
	env.Logger.Log("METHOD", "Get", "SPOT", "method start", "time_start", start)
	_, err := env.Common.GetUserClientFromToken(r)
	if err != nil {
		env.Common.ErrorResponseHelper(w, "4007", err.Error(), http.StatusBadRequest)
		return
	}
	pathParams := mux.Vars(r)
	planMessageId := pathParams["plan_message_id"]

	if err := env.Common.ValidateId(planMessageId, "plan_message_id"); err != nil {
		env.Common.ErrorResponseHelper(w, "4008", err.Error(), http.StatusBadRequest)
		return
	}
	planMessage, err := env.PlanMessageRepo.Get(planMessageId)
	if err != nil {
		env.Common.ErrorResponseHelper(w, "4009", lib.ERROR_UNEXPECTED, http.StatusInternalServerError)
		return
	}
	env.Logger.Log("METHOD", "Get", "SPOT", "method end", "time_spent", time.Since(start))
	env.Common.SuccessResponseHelper(w, planMessage, http.StatusOK)
}

func (env *PlanMessageEnv) Update(w http.ResponseWriter, r *http.Request) {

	start := time.Now()
	env.Logger.Log("METHOD", "Update", "SPOT", "method start", "time_start", start)
	_, err := env.Common.GetUserClientFromToken(r)
	if err != nil {
		env.Common.ErrorResponseHelper(w, "4010", err.Error(), http.StatusBadRequest)
		return
	}
	var pathParamConf map[string]string
	pathParamConf = make(map[string]string)
	pathParamConf["plan_message_id"] = ""

	var paramConf map[string]models.ParamConf
	paramConf = make(map[string]models.ParamConf)
	paramConf["message"] = models.ParamConf{Required: true, Type: lib.STRING_LARGE, EmptyAllowed: false}
	paramConf["action"] = models.ParamConf{Required: true, Type: lib.STRING_ACTION_NAME, EmptyAllowed: false}

	pathParamValue, paramMap, errCode, err := env.Common.ValidateInputParameters(r, paramConf, pathParamConf)
	if err != nil {
		env.Common.ErrorResponseHelper(w, "6"+errCode, err.Error(), http.StatusBadRequest)
		return
	}

	planMessageId := pathParamValue["plan_message_id"]

	err = env.PlanMessageRepo.Update(paramMap, planMessageId)
	if err != nil {
		env.Common.ErrorResponseHelper(w, "4011", lib.ERROR_UNEXPECTED, http.StatusInternalServerError)
		return
	}
	responsePlanMessageId := map[string]string{"plan_message_id": planMessageId}
	env.Logger.Log("METHOD", "Update", "SPOT", "method end", "time_spent", time.Since(start))
	env.Common.SuccessResponseHelper(w, responsePlanMessageId, http.StatusOK)
}

func (env *PlanMessageEnv) Delete(w http.ResponseWriter, r *http.Request) {

	start := time.Now()
	env.Logger.Log("METHOD", "Delete", "SPOT", "method start", "time_start", start)

	_, err := env.Common.GetUserClientFromToken(r)
	if err != nil {
		env.Common.ErrorResponseHelper(w, "4012", err.Error(), http.StatusBadRequest)
		return
	}
	pathParams := mux.Vars(r)
	planMessageId := pathParams["plan_message_id"]
	if err := env.Common.ValidateId(planMessageId, "plan_message_id"); err != nil {
		env.Common.ErrorResponseHelper(w, "4013", err.Error(), http.StatusBadRequest)
		return
	}

	planMessage := models.PlanMessage{}
	planMessage.Id = planMessageId

	err = env.PlanMessageRepo.Delete(planMessage)
	if err != nil {
		env.Common.ErrorResponseHelper(w, "4014", lib.ERROR_UNEXPECTED, http.StatusInternalServerError)
		return
	}

	responsePlanMessageId := map[string]string{"plan_message_id": planMessageId}
	env.Logger.Log("METHOD", "Delete", "SPOT", "method end", "time_spent", time.Since(start))
	env.Common.SuccessResponseHelper(w, responsePlanMessageId, http.StatusOK)
}

//VIM command to serializing error code :let @a=4001 | %s/\d\d\d\d/\=''.(@a+setreg('a',@a+1))/g
