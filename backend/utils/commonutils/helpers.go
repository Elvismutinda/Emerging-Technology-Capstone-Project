package commonutils

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"time"
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

func GetPaginationParams(page, pageSize string) (int, int) {
	var pageParam, pageSizeParam int
	// Parse and validate page and page_size
	if p, err := strconv.Atoi(page); err == nil && p > 0 {
		pageParam = p
	}
	if ps, err := strconv.Atoi(pageSize); err == nil && ps > 0 {
		pageSizeParam = ps
	}

	return pageParam, pageSizeParam
}

func ExtractDatesFromUrl(r *http.Request) (*time.Time, *time.Time, error) {
	// Extract start and end dates from URL query params
	startDateStr := r.URL.Query().Get("start_date")
	endDateStr := r.URL.Query().Get("end_date")

	// check if empty
	if startDateStr == "" {
		startDateStr = time.Now().Format("2006-01-02")
	}

	if endDateStr == "" {
		endDateStr = time.Now().Format("2006-01-02")
	}

	// Define the date format
	const dateFormat = "2006-01-02"

	// Parse the start date
	startDate, err := time.Parse(dateFormat, startDateStr)
	if err != nil {
		logrus.Errorf("Error parsing start date: %v", err)
		return nil, nil, errors.New("invalid start date format, expected YYYY-MM-DD")
	}

	// Parse the end date
	endDate, err := time.Parse(dateFormat, endDateStr)
	if err != nil {
		logrus.Errorf("Error parsing end date: %v", err)
		return nil, nil, errors.New("invalid end date format, expected YYYY-MM-DD")
	}

	// check if start date is less than end date
	if startDate.After(endDate) {
		return nil, nil, errors.New("start date must be after end date")
	}

	// Return the parsed dates
	return &startDate, &endDate, nil
}
