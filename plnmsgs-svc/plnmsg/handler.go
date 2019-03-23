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
		env.Common.ErrorResponseHelper(w, "7001", err.Error(), http.StatusBadRequest)
		return
	}

	var pathParamConf map[string]string
	pathParamConf = make(map[string]string)
	pathParamConf["plan_id"] = ""

	var paramConf map[string]models.ParamConf
	paramConf = make(map[string]models.ParamConf)
	paramConf["message"] = models.ParamConf{Required: true, Type: lib.STRING_LARGE, EmptyAllowed: false}
	paramConf["action"] = models.ParamConf{Required: false, Type: lib.STRING_SMALL, EmptyAllowed: false}

	pathParamValue, paramMap, errCode, err := env.Common.ValidateInputParameters(r, paramConf, pathParamConf)
	if err != nil {
		env.Common.ErrorResponseHelper(w, "6"+errCode, err.Error(), http.StatusBadRequest)
		return
	}

	planId := pathParamValue["plan_id"]

	planMessageId, err := env.PlanMessageRepo.Create(paramMap, planId)
	if err != nil {
		env.Common.ErrorResponseHelper(w, "7009", lib.ERROR_UNEXPECTED, http.StatusInternalServerError)
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
		env.Common.ErrorResponseHelper(w, "6010", err.Error(), http.StatusBadRequest)
		return
	}

	limit, offset, orderby, sort, err := env.Common.ValidateQueryString(r, "100", "0", "created_at", "asc")
	if err != nil {
		env.Common.ErrorResponseHelper(w, "6011", err.Error(), http.StatusBadRequest)
		return
	}
	planMessages, count, err := env.PlanMessageRepo.GetAll(limit, offset, orderby, sort)
	if err != nil {
		env.Common.ErrorResponseHelper(w, "6012", lib.ERROR_UNEXPECTED, http.StatusInternalServerError)
	}
	env.Logger.Log("METHOD", "GetAll", "SPOT", "method end", "time_spent", time.Since(start))
	env.Common.SuccessResponseList(w, planMessages, offset, limit, count)
}
func (env *PlanMessageEnv) Get(w http.ResponseWriter, r *http.Request) {

	start := time.Now()
	env.Logger.Log("METHOD", "Get", "SPOT", "method start", "time_start", start)
	_, err := env.Common.GetUserClientFromToken(r)
	if err != nil {
		env.Common.ErrorResponseHelper(w, "7013", err.Error(), http.StatusBadRequest)
		return
	}
	pathParams := mux.Vars(r)
	planMessageId := pathParams["plan_message_id"]

	if err := env.Common.ValidateId(planMessageId, "plan_message_id"); err != nil {
		env.Common.ErrorResponseHelper(w, "7014", err.Error(), http.StatusBadRequest)
		return
	}
	planMessage, err := env.PlanMessageRepo.Get(planMessageId)
	if err != nil {
		env.Common.ErrorResponseHelper(w, "7015", lib.ERROR_UNEXPECTED, http.StatusInternalServerError)
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
		env.Common.ErrorResponseHelper(w, "6016", err.Error(), http.StatusBadRequest)
		return
	}
	var pathParamConf map[string]string
	pathParamConf = make(map[string]string)
	pathParamConf["plan_message_id"] = ""

	var paramConf map[string]models.ParamConf
	paramConf = make(map[string]models.ParamConf)
	paramConf["message"] = models.ParamConf{Required: true, Type: lib.STRING_LARGE, EmptyAllowed: false}
	paramConf["action"] = models.ParamConf{Required: false, Type: lib.STRING_SMALL, EmptyAllowed: false}

	pathParamValue, paramMap, errCode, err := env.Common.ValidateInputParameters(r, paramConf, pathParamConf)
	if err != nil {
		env.Common.ErrorResponseHelper(w, "6"+errCode, err.Error(), http.StatusBadRequest)
		return
	}

	planMessageId := pathParamValue["plan_message_id"]

	err = env.PlanMessageRepo.Update(paramMap, planMessageId)
	if err != nil {
		env.Common.ErrorResponseHelper(w, "6025", lib.ERROR_UNEXPECTED, http.StatusInternalServerError)
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
		env.Common.ErrorResponseHelper(w, "6026", err.Error(), http.StatusBadRequest)
		return
	}
	pathParams := mux.Vars(r)
	planMessageId := pathParams["plan_message_id"]
	if err := env.Common.ValidateId(planMessageId, "plan_message_id"); err != nil {
		env.Common.ErrorResponseHelper(w, "6027", err.Error(), http.StatusBadRequest)
		return
	}

	planMessage := models.PlanMessage{}
	planMessage.Id = planMessageId

	err = env.PlanMessageRepo.Delete(planMessage)
	if err != nil {
		env.Common.ErrorResponseHelper(w, "6028", lib.ERROR_UNEXPECTED, http.StatusInternalServerError)
		return
	}

	responsePlanMessageId := map[string]string{"plan_message_id": planMessageId}
	env.Logger.Log("METHOD", "Delete", "SPOT", "method end", "time_spent", time.Since(start))
	env.Common.SuccessResponseHelper(w, responsePlanMessageId, http.StatusOK)
}
