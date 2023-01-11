package timechimp

import (
	"fmt"
	"net/http"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	t_types "github.com/leapforce-libraries/go_timechimp/types"
)

// User stores User from Service
//
type User struct {
	Id                     int64                    `json:"id"`
	UserName               *string                  `json:"userName"`
	DisplayName            *string                  `json:"displayName"`
	AccountType            int64                    `json:"accountType"`
	IsLocked               bool                     `json:"isLocked"`
	Picture                *string                  `json:"picture"`
	TagNames               []string                 `json:"tagNames"`
	Language               string                   `json:"language"`
	ContractHours          *float64                 `json:"contractHours"`
	ContractHourlyRate     *float64                 `json:"contractHourlyRate"`
	ContractCostHourlyRate *float64                 `json:"contractCostHourlyRate"`
	ContractStartDate      *t_types.DateTimeString  `json:"contractStartDate"`
	ContractEndDate        *t_types.DateTimeString  `json:"contractEndDate"`
	Created                t_types.DateTimeMsString `json:"created"`
	TeamName               *string                  `json:"teamName"`
	EmployeeNumber         *string                  `json:"employeeNumber"`
	Active                 bool                     `json:"active"`
}

// GetUsers returns all users
//
func (service *Service) GetUsers() (*[]User, *errortools.Error) {
	var users []User

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodGet,
		Url:           service.url("users"),
		ResponseModel: &users,
	}

	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &users, nil
}

// GetUsers returns all users
//
func (service *Service) GetUser(userId int64) (*User, *errortools.Error) {
	var user User

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodGet,
		Url:           service.url(fmt.Sprintf("users/%v", userId)),
		ResponseModel: &user,
	}

	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &user, nil
}
