package commonutils

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
}

func setResponseHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Request-Id", uuid.NewString())
	w.Header().Set("Content-Security-Policy", "default-src 'self';")
	w.Header().Set("Strict-Transport-Security", "max-age="+"10000000000000")
}

func HTTPResponse(w http.ResponseWriter, response Response, statusCode int) error {
	setResponseHeaders(w)
	w.WriteHeader(statusCode)
	// Write response
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		return err
	}

	return nil
}

func GetUserIdFromContext(ctx context.Context) (string, bool) {
	userID, ok := ctx.Value("userId").(string)
	return userID, ok
}

func CompareUserIds(urlUserId, userId string) (bool, error) {
	if urlUserId == "" {
		logrus.Error("User id cannot be empty")
		return false, errors.New("user id cannot be empty")
	}

	// check if the url param userId and the one in the context are the same
	if urlUserId != userId {
		logrus.Error("User id does not match")
		return false, errors.New("user id does not match")
	}
	return true, nil
}
