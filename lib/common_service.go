package lib

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"

	"github.com/syedomair/plan-api/models"

	jwt "github.com/dgrijalva/jwt-go"
	log "github.com/go-kit/kit/log"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/gorilla/mux"
)

type CommonService struct {
	response map[string]interface{}
	Logger   log.Logger
}

func (c CommonService) ErrorResponseHelper(w http.ResponseWriter, errorCode string, errorMessage string, httpStatus int) {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(httpStatus)
	w.Write(c.errorResponse(errorCode, errorMessage))
	c.Logger.Log("METHOD", "ErrorResponseHelper", "END", c.errorResponse(errorCode, errorMessage))
	return
}

func (c CommonService) SuccessResponseHelper(w http.ResponseWriter, class interface{}, httpStatus int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)
	w.Write(c.successResponse(class))
	c.Logger.Log("METHOD", "SuccessResponseHelper", "END", c.successResponse(class))
	return
}

func (c CommonService) SuccessResponseList(w http.ResponseWriter, class interface{}, offset string, limit string, count string) {
	tempResponse := make(map[string]interface{})
	tempResponse["offset"] = offset
	tempResponse["limit"] = limit
	tempResponse["count"] = count
	tempResponse["list"] = class

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(c.successResponse(tempResponse))
	c.Logger.Log("METHOD", "SuccessResponseList", "END", c.successResponse(tempResponse))
	return
}

func (c CommonService) errorResponse(errorCode string, message string) []byte {
	class := map[string]string{"error_code": errorCode, "message": message}
	return c.commonResponse(class, FAILURE)
}

func (c CommonService) successResponse(class interface{}) []byte {
	return c.commonResponse(class, SUCCESS)
}

func (c CommonService) commonResponse(class interface{}, result string) []byte {
	c.response = make(map[string]interface{})
	c.response["result"] = result
	c.response["data"] = class
	jsonByte, _ := json.Marshal(c.response)
	return jsonByte
}

func (c CommonService) ValidateId(id string, fieldName string) error {
	if id != "" {
		if err := validation.Validate(
			id,
			validation.Required.Error(fieldName+" is a required field"),
			is.UUIDv4.Error("invalid "+fieldName)); err != nil {
			return err
		}
	}
	return nil
}

func (c CommonService) checkAccessLevelValue(value interface{}) error {
	i, _ := strconv.Atoi(value.(string))
	if i < 0 || i > 10 {
		return errors.New("value must be from 0 and 10")
	}
	return nil
}
func (c CommonService) checkBooleanValue(value interface{}) error {
	i, _ := strconv.Atoi(value.(string))
	if i < 0 || i > 1 {
		return errors.New("value must be either 0 or 1")
	}
	return nil
}
func (c CommonService) checkActionNameValue(value interface{}) error {
	m := make(map[string]string)
	m["COST_UPDATE"] = "COST_UPDATE"
	m["VALIDITY_UPDATE"] = "VALIDITY_UPDATE"

	_, exists := m[value.(string)]
	if !exists {
		return errors.New("action value can only be COST_UPDATE, VALIDITY_UPDATE")
	}

	return nil
}

func (c CommonService) ValidateInputParameters(r *http.Request, paramConf map[string]models.ParamConf, pathParamConf map[string]string) (map[string]string, map[string]interface{}, string, error) {

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576)) // read the body of the request
	if err != nil {
		return nil, nil, "101", errors.New("Error reading the request body: " + err.Error())
	}
	var jsonMap map[string]interface{}
	decodedJson := json.NewDecoder(bytes.NewReader(body))

	if err := decodedJson.Decode(&jsonMap); err != nil {
		return nil, nil, "102", errors.New("Invalid JSON in request body: " + err.Error())
	}

	for k, _ := range jsonMap {
		if _, ok := paramConf[k]; !ok {
			return nil, nil, "103", errors.New(k + " Invalid parameter in JSON")
		}
	}

	for k, v := range paramConf {
		if val, ok := jsonMap[k]; ok {
			if reflect.TypeOf(val).String() != "string" {
				return nil, nil, "104", errors.New(k + " must be a valid string parameter")
			}
			if !v.EmptyAllowed {
				if err := validation.Validate(
					val,
					validation.Required.Error(k+" cannot be blank"),
				); err != nil {
					return nil, nil, "105", err
				}
			}
			switch fieldType := v.Type; fieldType {
			case STRING_SMALL:
				if err := validation.Validate(
					val,
					validation.Length(1, 32).Error(k+" allowed with max character of 32")); err != nil {
					return nil, nil, "106", err
				}

			case STRING_LARGE:
				if err := validation.Validate(
					val,
					validation.Length(1, 1000).Error(k+" allowed with max character of 1000")); err != nil {
					return nil, nil, "107", err
				}

			case INT:
				if err := validation.Validate(
					val,
					is.Int.Error(k+" must be a valid integer")); err != nil {
					return nil, nil, "108", err
				}
			case INT_ACCESS_LEVEL:
				if err := validation.Validate(
					val,
					is.Int.Error(k+" must be a valid integer")); err != nil {
					return nil, nil, "109", err
				} else if err := validation.Validate(
					val,
					validation.By(c.checkAccessLevelValue)); err != nil {
					return nil, nil, "110", errors.New(k + " " + err.Error())
				}
			case INT_BOOLEAN:
				if err := validation.Validate(
					val,
					is.Int.Error(k+" must be a valid integer")); err != nil {
					return nil, nil, "111", err
				} else if err := validation.Validate(
					val,
					validation.By(c.checkBooleanValue)); err != nil {
					return nil, nil, "112", errors.New(k + " " + err.Error())
				}
			case ID:
				if val != "00000000-0000-0000-0000-000000000000" {
					if err := validation.Validate(
						val,
						is.UUIDv4.Error(k+" invalid ID")); err != nil {
						return nil, nil, "113", err
					}
				}
			case STRING_EMAIL:
				if err := validation.Validate(
					val,
					is.Email.Error(k+" must be a valid email address")); err != nil {
					return nil, nil, "114", err
				}
			case STRING_PASSWORD:
				if err := validation.Validate(
					val,
					validation.Length(6, 32).Error(k+" must be atleast 5 characters long")); err != nil {
					return nil, nil, "115", err
				}
			case STRING_NAME:
				if err := validation.Validate(
					val,
					validation.Length(1, 50).Error(k+" field size cannot be  greater than 50 characters")); err != nil {
					return nil, nil, "116", err
				}
			case STRING_ACTION_NAME:
				if err := validation.Validate(
					val,
					validation.By(c.checkActionNameValue)); err != nil {
					return nil, nil, "117", errors.New(k + " " + err.Error())
				}
			}

		} else {
			if v.Required {
				return nil, nil, "118", errors.New(k + " is a requird field")
			}
		}

	}

	pathParams := mux.Vars(r)
	for k, _ := range pathParamConf {
		tempId := pathParams[k]

		if err := validation.Validate(
			tempId,
			is.UUIDv4.Error(k+" invalid ID")); err != nil {
			return nil, nil, "119", err
		}
		pathParamConf[k] = tempId

	}

	return pathParamConf, jsonMap, "", nil
}

func (c CommonService) ValidateQueryString(r *http.Request, defaultLimit string, defaultOffset string, defaultOrderby string, defaultSort string) (string, string, string, string, error) {

	limit := ""
	offset := ""
	orderby := ""
	sort := ""

	limits, ok := r.URL.Query()["limit"]
	if !ok || len(limits[0]) < 1 {
		limit = ""
	} else {
		limit = limits[0]
	}

	offsets, ok := r.URL.Query()["offset"]
	if !ok || len(offsets[0]) < 1 {
		offset = ""
	} else {
		offset = offsets[0]
	}
	orderbys, ok := r.URL.Query()["orderby"]
	if !ok || len(orderbys[0]) < 1 {
		orderby = ""
	} else {
		orderby = orderbys[0]
	}
	sorts, ok := r.URL.Query()["sort"]
	if !ok || len(sorts[0]) < 1 {
		sort = ""
	} else {
		sort = sorts[0]
	}

	if limit != "" {
		if _, err := strconv.Atoi(limit); err != nil {
			return "", "", "", "", errors.New("Invalid 'limit' number in query string. Must be a number. ")
		}
	} else {
		limit = defaultLimit
	}
	if offset != "" {
		if _, err := strconv.Atoi(offset); err != nil {
			return "", "", "", "", errors.New("Invalid 'offset' number in query string. Must be a number. ")
		}
	} else {
		offset = defaultOffset
	}

	if orderby != "" {
		if _, err := strconv.Atoi(orderby); err == nil {
			return "", "", "", "", errors.New("Invalid 'orderby' value in query string. Must be a string. ")
		}
	} else {
		orderby = defaultOrderby
	}

	if sort != "" {
		if _, err := strconv.Atoi(sort); err == nil {
			return "", "", "", "", errors.New("Invalid 'sort' value in query string. Must be a string. ")
		}
		if (sort != "asc") && (sort != "desc") {
			return "", "", "", "", errors.New("Invalid 'sort' value in query string. Must be either 'asc' or 'desc'. ")
		}
	} else {
		sort = defaultSort
	}

	return limit, offset, orderby, sort, nil
}

func (c CommonService) GetUserClientFromToken(r *http.Request) (string, error) {
	tokenString := r.Header.Get("Token")
	if tokenString == "" {
		return "", errors.New("Missing Token in header")
	}

	type Claims struct {
		CurrentUserId string `json:"current_user_id"`
		jwt.StandardClaims
	}
	tokenClaims := Claims{}

	token, err := jwt.ParseWithClaims(tokenString, &tokenClaims, func(token *jwt.Token) (interface{}, error) {
		return []byte(SIGNING_KEY), nil
	})
	if err != nil {
		return "", errors.New(err.Error())
	}
	if token.Valid {
		return tokenClaims.CurrentUserId, nil
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return "", errors.New("That's not even a token")
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			// Token is either expired or not active yet
			return "", errors.New("Timing is everything")
		} else {
			return "", errors.New("Couldn't handle this token")
		}
	} else {
		return "", errors.New("Couldn't handle this token")
	}
}
