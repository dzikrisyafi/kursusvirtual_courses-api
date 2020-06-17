package rest

import (
	"errors"
	"fmt"
	"time"

	"github.com/dzikrisyafi/kursusvirtual_utils-go/rest_errors"
	"github.com/golang-restclient/rest"
)

var (
	QuizRepository quizRepositoryInterface = &quizRepository{}
	quizRestClient                         = rest.RequestBuilder{
		BaseURL: "http://localhost:8004",
		Timeout: 100 * time.Millisecond,
	}
)

type quizRepository struct{}

type quizRepositoryInterface interface {
	DeleteQuiz(int) rest_errors.RestErr
}

func (r *quizRepository) DeleteQuiz(courseID int) rest_errors.RestErr {
	response := quizRestClient.Delete(fmt.Sprintf("/internal/quiz/%d", courseID))

	if response == nil || response.Response == nil {
		return rest_errors.NewInternalServerError("invalid rest client response when trying to delete quiz", errors.New("rest client error"))
	}

	if response.StatusCode > 299 {
		apiErr, err := rest_errors.NewRestErrorFromBytes(response.Bytes())
		if err != nil {
			return rest_errors.NewInternalServerError("invalid error interface when trying to delete quiz", err)
		}

		return apiErr
	}

	return nil
}
