package utils

import (
	"github.com/apsdehal/go-logger"
	"net/http"
	"os"
	"time"
)

// http configuration
const (
	ReadWriteTimeout = 100
	TimeOutDuration  = 5 * time.Second
	IdleTimeout      = 10
)

// time format
const TimeFormat = "2006-01-02 15:04:05"

var StatusMessage = map[string]string{
	"successfully_seed_mock": "Successfully seed example data",
	"live":                   "Api is Live",
	"invalid_json":           "Invalid payload",
	"status_ok":              http.StatusText(http.StatusOK),
	"status_deleted":         "Successfully deleted",
	"successfully_created":   "Successfully created",
	"failed_created":         "Failed Created",
	"failed_updated":         "Failed Updated",
	"failed_deleted":         "Failed Deleted",
}

var HttpMessage = map[string]map[string]interface{}{
	"successfully_seed_mock": {"status": http.StatusOK, "message": StatusMessage["successfully_seed_mock"]},
	"live":                   {"status": http.StatusOK, "message": StatusMessage["live"]},
	"invalid_json":           {"status": http.StatusBadRequest, "message": StatusMessage["invalid_json"]},
	"status_ok":              {"status": http.StatusOK, "message": StatusMessage["status_ok"]},
	"status_deleted":         {"status": http.StatusNoContent, "message": StatusMessage["status_deleted"]},
	"successfully_created":   {"status": http.StatusCreated, "message": StatusMessage["successfully_created"]},
	"failed_created":         {"status": http.StatusCreated, "message": StatusMessage["failed_created"]},
	"failed_updated":         {"status": http.StatusCreated, "message": StatusMessage["failed_updated"]},
	"failed_deleted":         {"status": http.StatusCreated, "message": StatusMessage["failed_deleted"]},
}

// logging configuration
var ErrorLog, _ = logger.New("[ERROR] - ", 5, os.Stdout)
var DebugLog, _ = logger.New("[DEBUG] - ", 2, os.Stdout)
var InfoLog, _ = logger.New("[INFO] - ", 6, os.Stdout)

// example api for example mocking data
// from https://jsonplaceholder.typicode.com/
var ExampleJsonData = map[string]string{
	"users":    "https://jsonplaceholder.typicode.com/users",
	"todos":    "https://jsonplaceholder.typicode.com/todos",
	"photos":   "https://jsonplaceholder.typicode.com/photos",
	"albums":   "https://jsonplaceholder.typicode.com/albums",
	"comments": "https://jsonplaceholder.typicode.com/comments",
	"posts":    "https://jsonplaceholder.typicode.com/posts",
}
