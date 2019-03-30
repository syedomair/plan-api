package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/syedomair/plan-api/lib"
	"github.com/syedomair/plan-api/test/testdata"
)

type testCaseType struct {
	method         string
	url            string
	path           string
	pathParam      string
	header         string
	requestBody    string
	responseResult string
	responseData   string
}

var testCases []testCaseType

func main() {
	arg := ""
	if len(os.Args) > 1 {
		arg = os.Args[1]
		fmt.Println("Command line argument: ", arg)
	}

	url := "https://plans-api.herokuapp.com/"
	if arg == "local" {
		url = "http://localhost:8180/"
	}
	//VIM command for numbering comments
	//:let @a=100 | %s/num\d\d\d/\='num'.(@a+setreg('a',@a+1))/g
	currentTime := time.Now()
	uniqueString := strconv.FormatInt(currentTime.UnixNano(), 10)
	uniqueEmail := "email_" + uniqueString + "@gmail.com"

	testCases = []testCaseType{
		//num100 missing last_name field
		{"POST", url, "register", "", "", `{"first_name":"` + testdata.ValidFirstName + `", "email":"` + uniqueEmail + `", "password":"` + testdata.ValidPassword + `" }`, lib.FAILURE, ``},
		//num101 last_name length
		{"POST", url, "register", "", "", `{"first_name":"` + testdata.ValidFirstName + `", "last_name":"` + testdata.InValidLenLastName + `", "email":"` + uniqueEmail + `", "password":"` + testdata.ValidPassword + `"}`, lib.FAILURE, ``},
		//num102 last_name invalid
		{"POST", url, "register", "", "", `{"first_name":"` + testdata.ValidFirstName + `", "last_name":"` + testdata.InValidLastName + `", "email":"` + uniqueEmail + `", "password":"` + testdata.ValidPassword + `"}`, lib.FAILURE, ``},
		//num103 missing first_name
		{"POST", url, "register", "", "", `{"last_name":"` + testdata.ValidLastName + `", "email":"` + uniqueEmail + `", "password":"` + testdata.ValidPassword + `"}`, lib.FAILURE, ``},
		//num104 first_name length
		{"POST", url, "register", "", "", `{"first_name":"` + testdata.InValidLenFirstName + `", "last_name":"` + testdata.ValidLastName + `", "email":"` + uniqueEmail + `", "password":"` + testdata.ValidPassword + `"}`, lib.FAILURE, ``},
		//num105 first_name invalid
		{"POST", url, "register", "", "", `{"first_name":"` + testdata.InValidFirstName + `", "last_name":"` + testdata.ValidLastName + `", "email":"` + uniqueEmail + `", "password":"` + testdata.ValidPassword + `"}`, lib.FAILURE, ``},
		//num106 all good success
		{"POST", url, "register", "", "", `{"first_name":"` + testdata.ValidFirstName + `", "last_name":"` + testdata.ValidLastName + `", "email":"` + uniqueEmail + `", "password":"` + testdata.ValidPassword + `"}`, lib.SUCCESS, ``},
		//num107 In valid email
		{"POST", url, "register", "", "", `{"first_name":"` + testdata.ValidFirstName + `", "last_name":"` + testdata.ValidLastName + `", "email":"` + testdata.InValidEmail + `", "password":"` + testdata.ValidPassword + `"}`, lib.FAILURE, ``},
		//num108 email blank
		{"POST", url, "register", "", "", `{"first_name":"` + testdata.ValidFirstName + `", "last_name":"` + testdata.ValidLastName + `", "email":"` + testdata.InValidEmailBlank + `", "password":"` + testdata.ValidPassword + `"}`, lib.FAILURE, ``},
		//num109 email missing
		{"POST", url, "register", "", "", `{"first_name":"` + testdata.ValidFirstName + `", "last_name":"` + testdata.ValidLastName + `", "password":"` + testdata.ValidPassword + `"}`, lib.FAILURE, ``},
		//num110 unique email address test
		{"POST", url, "register", "", "", `{"first_name":"` + testdata.ValidFirstName + `", "last_name":"` + testdata.ValidLastName + `", "email":"` + uniqueEmail + `", "password":"` + testdata.ValidPassword + `"}`, lib.FAILURE, ``},
		//num111 missing password
		{"POST", url, "register", "", "", `{"first_name":"` + testdata.ValidFirstName + `", "last_name":"` + testdata.ValidLastName + `", "email":"` + uniqueEmail + `"}`, lib.FAILURE, ``},
		//num112 password length
		{"POST", url, "register", "", "", `{"first_name":"` + testdata.ValidFirstName + `", "last_name":"` + testdata.ValidLastName + `", "email":"` + uniqueEmail + `", "password":"` + testdata.InValidPassword + `"}`, lib.FAILURE, ``},
		//num113 password blank
		{"POST", url, "register", "", "", `{"first_name":"` + testdata.ValidFirstName + `", "last_name":"` + testdata.ValidLastName + `", "email":"` + uniqueEmail + `", "password":"` + testdata.InValidPasswordBlank + `"}`, lib.FAILURE, ``},
		//num114 same email exists
		{"POST", url, "register", "", "", `{"first_name":"` + testdata.ValidFirstName + `", "last_name":"` + testdata.ValidLastName + `", "email":"` + uniqueEmail + `", "password":"` + testdata.ValidPassword + `"}`, lib.FAILURE, ``},
		//num115 all good
		{"POST", url, "register", "", "", `{"first_name":"` + testdata.ValidFirstName + `", "last_name":"` + testdata.ValidLastName + `", "email":"` + uniqueEmail + `.eu", "password":"` + testdata.ValidPassword + `"}`, lib.SUCCESS, ``},
		//num116 all good
		{"POST", url, "login", "", "", `{"email":"` + uniqueEmail + `.eu", "password":"` + testdata.ValidPassword + `"}`, lib.SUCCESS, ``},
		//num117 Login  ... returns token
		{"POST", url, "login", "", "", `{"email":"` + testdata.ValidEmail + `", "password":"` + testdata.ValidPassword + `"}`, lib.SUCCESS, ``},
		//num118 Login  ...returns token
		{"POST", url, "login", "", "", `{"email":"` + uniqueEmail + `.eu", "password":"` + testdata.ValidPassword + `"}`, lib.SUCCESS, ``},
		//num119 GET ALL User
		{"GET", url, "users", "", "token", ``, lib.SUCCESS, ``},
		//num120 Plan invalid title
		{"POST", url, "plans", "", "token", `{"title":"` + testdata.InValidLenPlanTitle + `" , "status":"` + testdata.ValidPlanStatus + `"  , "cost":"` + testdata.ValidPlanCost + `", "validity":"` + testdata.ValidPlanValidity + `" }`, lib.FAILURE, ``},
		//num121 Plan invalid title blank
		{"POST", url, "plans", "", "token", `{"title":"` + testdata.InValidPlanTitleBlank + `" , "status":"` + testdata.ValidPlanStatus + `"  , "cost":"` + testdata.ValidPlanCost + `", "validity":"` + testdata.ValidPlanValidity + `" }`, lib.FAILURE, ``},
		//num122 Plan invalid cost
		{"POST", url, "plans", "", "token", `{"title":"` + testdata.ValidPlanTitle + `" , "status":"` + testdata.ValidPlanStatus + `"  , "cost":"` + testdata.InValidPlanCost + `", "validity":"` + testdata.ValidPlanValidity + `" }`, lib.FAILURE, ``},
		//num123 Plan invalid cost blank
		{"POST", url, "plans", "", "token", `{"title":"` + testdata.ValidPlanTitle + `" , "status":"` + testdata.ValidPlanStatus + `"  , "cost":"` + testdata.InValidPlanCostBlank + `", "validity":"` + testdata.ValidPlanValidity + `" }`, lib.FAILURE, ``},
		//num124 Plan invalid validity
		{"POST", url, "plans", "", "token", `{"title":"` + testdata.ValidPlanTitle + `" , "status":"` + testdata.ValidPlanStatus + `"  , "cost":"` + testdata.ValidPlanCost + `", "validity":"` + testdata.InValidPlanValidity + `" }`, lib.FAILURE, ``},
		//num125 Plan invalid validity blank
		{"POST", url, "plans", "", "token", `{"title":"` + testdata.ValidPlanTitle + `" , "status":"` + testdata.ValidPlanStatus + `"  , "cost":"` + testdata.InValidPlanCostBlank + `", "validity":"` + testdata.InValidPlanValidityBlank + `" }`, lib.FAILURE, ``},
		//num126 Plan invalid status
		{"POST", url, "plans", "", "token", `{"title":"` + testdata.ValidPlanTitle + `" , "status":"` + testdata.InValidPlanStatus + `"  , "cost":"` + testdata.ValidPlanCost + `", "validity":"` + testdata.ValidPlanValidity + `" }`, lib.FAILURE, ``},
		//num127 Plan invalid status blank
		{"POST", url, "plans", "", "token", `{"title":"` + testdata.ValidPlanTitle + `" , "status":"` + testdata.InValidPlanStatusBlank + `"  , "cost":"` + testdata.ValidPlanCost + `", "validity":"` + testdata.ValidPlanValidity + `" }`, lib.FAILURE, ``},
		//num128 Plan all good
		{"POST", url, "plans", "", "token", `{"title":"` + testdata.ValidPlanTitle + `" , "status":"` + testdata.ValidPlanStatus + `"  , "cost":"` + testdata.ValidPlanCost + `", "validity":"` + testdata.ValidPlanValidity + `" }`, lib.SUCCESS, ``},
		//num129 Update invalid cost
		{"PATCH", url, "plans", "plan_id", "token", `{"title":"` + testdata.ValidPlanTitle + `" , "status":"` + testdata.ValidPlanStatus + `"  , "cost":"` + testdata.InValidPlanCost + `", "validity":"` + testdata.ValidPlanValidity + `" }`, lib.FAILURE, ``},
		//num130 Update valid cost
		{"PATCH", url, "plans", "plan_id", "token", `{"title":"` + testdata.ValidPlanTitle + `" , "status":"` + testdata.ValidPlanStatus + `"  , "cost":"` + testdata.ValidPlanCost + `", "validity":"` + testdata.ValidPlanValidity + `" }`, lib.SUCCESS, ``},
		//num131 GET ALL
		{"GET", url, "plans", "", "token", ``, lib.SUCCESS, ``},
		//num132 GET
		{"GET", url, "plans", "plan_id", "token", ``, lib.SUCCESS, ``},
		//num133 plan-messages invalid message
		{"POST", url, "plan-messages", "plan_id", "token", `{"message":"` + testdata.InValidPlanMessageBlank + `" , "action":"` + testdata.ValidPlanMessageAction + `" }`, lib.FAILURE, ``},
		//num134 plan-messages invalid action
		{"POST", url, "plan-messages", "plan_id", "token", `{"message":"` + testdata.ValidPlanMessage + `" , "action":"` + testdata.InValidPlanMessageActionBlank + `" }`, lib.FAILURE, ``},
		//num135 plan-messages all good
		{"POST", url, "plan-messages", "plan_id", "token", `{"message":"` + testdata.ValidPlanMessage + `" , "action":"` + testdata.ValidPlanMessageAction + `" }`, lib.SUCCESS, ``},
		//num136 plan-messages Update
		{"PATCH", url, "plan-messages", "plan_message_id", "token", `{"message":"` + testdata.ValidPlanMessage + `" , "action":"` + testdata.ValidPlanMessageAction + `" }`, lib.SUCCESS, ``},
		//num137 plan-messages Get
		{"GET", url, "plan-messages", "plan_message_id", "token", ``, lib.SUCCESS, ``},
		//num138 plan-messages Get ALL
		{"GET", url, "plan-messages", "plan_id", "token", ``, lib.SUCCESS, ``},
		//num139 plan-messages delete
		{"DELETE", url, "plan-messages", "plan_message_id", "token", ``, lib.SUCCESS, ``},
		//num140 plan delete
		{"DELETE", url, "plans", "plan_id", "token", ``, lib.SUCCESS, ``},
		//num141 stats total user count
		{"GET", url, "stats/user-count", "", "token", ``, lib.SUCCESS, ``},
		//num142 stats user registration data
		{"GET", url, "stats/user-reg-data", "", "token", ``, lib.SUCCESS, ``},
		//num143 stats plan data
		{"GET", url, "stats/plan-data", "", "token", ``, lib.SUCCESS, ``},
		/*
		 */
	}
	i := 100
	token := ""
	planId := ""
	planMessageId := ""
	for _, testCase := range testCases {
		url := ""
		if testCase.pathParam == "plan_id" {
			url = testCase.url + testCase.path + "/" + planId
		} else if testCase.pathParam == "plan_message_id" {
			url = testCase.url + testCase.path + "/" + planMessageId
		} else {
			url = testCase.url + testCase.path
		}
		fmt.Println(url)
		req, err := http.NewRequest(testCase.method, url, strings.NewReader(testCase.requestBody))
		if err != nil {
			print(err)
		}

		if testCase.header == "token" {
			req.Header.Set("Token", token)
		}

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			print(err)
		}

		body, _ := ioutil.ReadAll(resp.Body)

		var bodyInterface map[string]interface{}

		json.Unmarshal(body, &bodyInterface)
		jsonResult, _ := json.Marshal(bodyInterface["result"])
		jsonData, _ := json.Marshal(bodyInterface["data"])
		fmt.Println(string(body))

		if strings.Compare(strings.Trim(string(jsonResult), "\""), testCase.responseResult) == 0 {
			fmt.Println("\033[32m" + strconv.Itoa(i) + " PASS" + "\033[39m")
		} else {
			fmt.Println("\033[31m" + strconv.Itoa(i) + " FAIL" + "\033[39m")
			fmt.Println(strconv.Itoa(i) + " " + testCase.method + " " + string(testCase.url) + " " + testCase.path + " " + testCase.requestBody)
			fmt.Println(string(jsonData))
			fmt.Println(string(jsonResult))
		}
		if testCase.method == "POST" && testCase.path == "login" && testCase.responseResult == lib.SUCCESS {
			var tokenInterface map[string]interface{}
			json.Unmarshal(jsonData, &tokenInterface)
			jsonToken, _ := json.Marshal(tokenInterface["token"])
			token = string(bytes.Trim(jsonToken, `"`))
			fmt.Println("token:", token)
		}
		if testCase.method == "POST" && testCase.path == "plans" && testCase.responseResult == lib.SUCCESS {
			var planIdInterface map[string]interface{}
			json.Unmarshal(jsonData, &planIdInterface)
			jsonPlanId, _ := json.Marshal(planIdInterface["plan_id"])
			planId = string(bytes.Trim(jsonPlanId, `"`))
			fmt.Println("planId:", planId)
		}
		if testCase.method == "POST" && testCase.path == "plan-messages" && testCase.responseResult == lib.SUCCESS {
			var planMessageIdInterface map[string]interface{}
			json.Unmarshal(jsonData, &planMessageIdInterface)
			jsonPlanMessageId, _ := json.Marshal(planMessageIdInterface["plan_message_id"])
			planMessageId = string(bytes.Trim(jsonPlanMessageId, `"`))
			fmt.Println("planMessageId:", planMessageId)
		}

		fmt.Println("---------------------------------------------------------------------------")

		i++
	}
}

func minikubeServiceURL(serviceName string) (string, error) {
	minikube := "minikube"
	service := "service"
	url := "--url"

	serviceURL, err := exec.Command(minikube, service, serviceName, url).Output()
	if err != nil {
		return "", err
	}
	return strings.Trim(string(serviceURL), "\n"), nil
}
