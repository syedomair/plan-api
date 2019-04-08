package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	goKitLog "github.com/go-kit/kit/log"

	"github.com/gorilla/handlers"
	lib "github.com/syedomair/plan-api/lib"
	plan "github.com/syedomair/plan-api/services/plans-svc/plan"
	planMessage "github.com/syedomair/plan-api/services/plnmsgs-svc/plnmsg"
	public "github.com/syedomair/plan-api/services/public-svc/public"
	stat "github.com/syedomair/plan-api/services/stats-svc/stat"
	user "github.com/syedomair/plan-api/services/users-svc/user"
)

type Env struct {
	Logger         goKitLog.Logger
	Common         lib.CommonService
	PublicEnv      public.PublicEnv
	PlanEnv        plan.PlanEnv
	PlanMessageEnv planMessage.PlanMessageEnv
	UserEnv        user.UserEnv
	StatEnv        stat.StatEnv
}

func main() {
	db, err := lib.CreateDBConnection()
	db.LogMode(true)
	defer db.Close()

	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	} else {
		fmt.Println("Connected to Database")
	}
	fmt.Println("application running on port: " + os.Getenv("PORT"))

	chJob := make(chan string, 100)

	logger := lib.GetLogger()
	statRepository := &stat.StatRepository{db, logger}
	publicRepository := &public.PublicRepository{db, logger}
	userRepository := &user.UserRepository{db, logger}
	planRepository := &plan.PlanRepository{db, logger, chJob}
	PlanMessageRepository := &planMessage.PlanMessageRepository{db, logger}
	commonService := lib.CommonService{Logger: logger}

	publicEnv := &public.PublicEnv{logger, publicRepository, commonService}
	userEnv := &user.UserEnv{logger, userRepository, commonService}
	statEnv := &stat.StatEnv{logger, statRepository, commonService}
	planEnv := &plan.PlanEnv{logger, planRepository, commonService}
	planMessageEnv := &planMessage.PlanMessageEnv{logger, PlanMessageRepository, commonService}

	env := &Env{PublicEnv: *publicEnv,
		UserEnv:        *userEnv,
		StatEnv:        *statEnv,
		PlanEnv:        *planEnv,
		PlanMessageEnv: *planMessageEnv,
		Common:         commonService,
		Logger:         logger}

	router := env.NewRouter()

	headersOk := handlers.AllowedHeaders([]string{"Apikey", "Token", "Content-Type", "Origin", "Accept", "Access-Control-Allow-Origin"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "PATCH", "DELETE"})

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}
