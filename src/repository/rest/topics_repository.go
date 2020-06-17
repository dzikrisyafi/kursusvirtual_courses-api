package rest

import (
	"errors"
	"fmt"
	"time"

	"github.com/dzikrisyafi/kursusvirtual_utils-go/rest_errors"
	"github.com/golang-restclient/rest"
)

var (
	TopicsRepository topicsRepositoryInterface = &topicsRepository{}
	topicsRestClient                           = rest.RequestBuilder{
		BaseURL: "http://localhost:8003",
		Timeout: 100 * time.Millisecond,
	}
)

type topicsRepository struct{}

type topicsRepositoryInterface interface {
	DeleteTopics(int) rest_errors.RestErr
}

func (r *topicsRepository) DeleteTopics(courseID int) rest_errors.RestErr {
	response := topicsRestClient.Delete(fmt.Sprintf("/internal/sections/%d", courseID))

	if response == nil || response.Response == nil {
		return rest_errors.NewInternalServerError("invalid rest client response when trying to delete topics", errors.New("rest client error"))
	}

	if response.StatusCode > 299 {
		apiErr, err := rest_errors.NewRestErrorFromBytes(response.Bytes())
		if err != nil {
			return rest_errors.NewInternalServerError("invalid error interface when trying to delete topics", err)
		}

		return apiErr
	}

	return nil
}
