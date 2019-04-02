package main

import (
	"net/http"
	"net/http/httputil"
	"time"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
	Access      string
}

type Routes []Route

func (env *Env) NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		env.Common.ErrorResponseHelper(w, "0001", "API Resource Not Found ", http.StatusNotFound)
	})

	var routes = Routes{
		Route{
			"Ping",
			"GET",
			"/ping",
			env.PublicEnv.Ping,
			"public",
		},
		Route{
			"Register",
			"POST",
			"/register",
			env.PublicEnv.Register,
			"public",
		},
		Route{
			"Login",
			"POST",
			"/login",
			env.PublicEnv.Login,
			"public",
		},
		//User
		Route{
			"UserGetAll",
			"GET",
			"/users",
			env.UserEnv.GetAll,
			"admin",
		},
		//Stat
		Route{
			"GetTotalUserCount,",
			"GET",
			"/stats/user-count",
			env.StatEnv.GetTotalUserCount,
			"admin",
		},
		Route{
			"GetTotalUserCountLast30Days,",
			"GET",
			"/stats/user-count-30-days",
			env.StatEnv.GetTotalUserCountLast30Days,
			"admin",
		},
		Route{
			"GetUserRegData,",
			"GET",
			"/stats/user-reg-data",
			env.StatEnv.GetUserRegData,
			"admin",
		},
		Route{
			"GetPlanData,",
			"GET",
			"/stats/plan-data",
			env.StatEnv.GetPlanData,
			"admin",
		},
		//Plan
		Route{
			"CreatePlan",
			"POST",
			"/plans",
			env.PlanEnv.CreatePlan,
			"user",
		},
		Route{
			"GetAll",
			"GET",
			"/plans",
			env.PlanEnv.GetAll,
			"user",
		},
		Route{
			"Get",
			"GET",
			"/plans/{plan_id}",
			env.PlanEnv.Get,
			"user",
		},
		Route{
			"Update",
			"PATCH",
			"/plans/{plan_id}",
			env.PlanEnv.Update,
			"user",
		},
		Route{
			"Delete",
			"DELETE",
			"/plans/{plan_id}",
			env.PlanEnv.Delete,
			"user",
		},
		//PlanMessage
		Route{
			"Create",
			"POST",
			"/plan-messages/{plan_id}",
			env.PlanMessageEnv.Create,
			"user",
		},
		Route{
			"GetAll",
			"GET",
			"/plan-messages/plan/{plan_id}",
			env.PlanMessageEnv.GetAll,
			"user",
		},
		Route{
			"Get",
			"GET",
			"/plan-messages/{plan_message_id}",
			env.PlanMessageEnv.Get,
			"user",
		},
		Route{
			"Update",
			"PATCH",
			"/plan-messages/{plan_message_id}",
			env.PlanMessageEnv.Update,
			"user",
		},
		Route{
			"Delete",
			"DELETE",
			"/plan-messages/{plan_message_id}",
			env.PlanMessageEnv.Delete,
			"user",
		},
	}
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = env.loggerMiddleware(handler)
		if route.Access == "user" {
			handler = env.securityUserMiddleware(handler)
		}

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}

	return router
}

func (env *Env) securityUserMiddleware(handle http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		env.Logger.Log("METHOD", "securityUserMiddleware", "SPOT", "METHOD START")
		start := time.Now()

		token := r.Header.Get("Token")
		env.Logger.Log("METHOD", "securityUserMiddleware", "SPOT", "TOKEN", "token", token)
		_, err := env.PublicEnv.PublicRepo.FindToken(token)

		if err != nil {
			env.Logger.Log("METHOD", "securityUserMiddleware", "SPOT", "TOKEN", "messeage", "Token not found")
			env.Common.ErrorResponseHelper(w, "0002", "Token not found, Please login again", http.StatusUnauthorized)
			return
		} else {
			handle.ServeHTTP(w, r)

		}
		env.Logger.Log("METHOD", "securityUserMiddleware", "SPOT", "METHOD END", "time_spent", time.Since(start))
	})
}

func (env *Env) loggerMiddleware(handle http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		env.Logger.Log("METHOD", "loggerMiddleware", "SPOT", "METHOD START")
		start := time.Now()

		requestDump, err := httputil.DumpRequest(r, true)
		if err != nil {
			env.Logger.Log("METHOD", "loggerMiddleware", "DumpRequest ERR:", err.Error())
		}
		env.Logger.Log("METHOD", "loggerMiddleware", "REQUEST:", string(requestDump))
		handle.ServeHTTP(w, r)

		env.Logger.Log("METHOD", "loggerMiddleware", "SPOT", "METHOD END", "time_spent", time.Since(start))
	})
}
