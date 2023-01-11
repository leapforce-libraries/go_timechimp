package timechimp

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	t_types "github.com/leapforce-libraries/go_timechimp/types"
)

// TimeEntry stores TimeEntry from Service
//
type TimeEntry struct {
	Id              int64                   `json:"id"`
	CustomerId      int64                   `json:"customerId"`
	CustomerName    *string                 `json:"customerName"`
	ProjectId       int64                   `json:"projectId"`
	ProjectName     *string                 `json:"projectName"`
	ProjectTaskId   int64                   `json:"projectTaskId"`
	TaskId          int64                   `json:"taskId"`
	TaskName        *string                 `json:"taskName"`
	UserId          int64                   `json:"userId"`
	UserDisplayName string                  `json:"userDisplayName"`
	UserTags        []string                `json:"userTags"`
	Date            t_types.DateTimeString  `json:"date"`
	Hours           float64                 `json:"hours"`
	Notes           *string                 `json:"notes"`
	StartEnd        *string                 `json:"startEnd"`
	Start           *t_types.DateTimeString `json:"start"`
	End             *t_types.DateTimeString `json:"end"`
	Pause           *float64                `json:"pause"`
	ExternalName    *string                 `json:"externalName"`
	ExternalURL     *string                 `json:"externalUrl"`
	Status          int64                   `json:"status"`
	StatusIntern    int64                   `json:"statusIntern"`
	StatusExtern    int64                   `json:"statusExtern"`
	Tags            []Tag                   `json:"tags"`
	Modified        t_types.DateTimeString  `json:"modified"`
}

type GetTimeEntriesConfig struct {
	Modified *time.Time
}

// GetTimeEntries returns all timeEntries
//
func (service *Service) GetTimeEntries(config *GetTimeEntriesConfig) (*[]TimeEntry, *errortools.Error) {
	values := url.Values{}
	endpoint := "time"

	if config != nil {
		if config.Modified != nil {
			values.Set("modified", t_types.DateTimeString(*config.Modified).String())
		}
	}

	path := fmt.Sprintf("%s?%s", endpoint, values.Encode())
	timeEntries := []TimeEntry{}

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodGet,
		Url:           service.url(path),
		ResponseModel: &timeEntries,
	}

	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &timeEntries, nil
}
