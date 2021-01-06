package utils

import (
	"github.com/google/uuid"
	"github.com/valyala/fasthttp"
	"strings"
	"time"
)

// GenerateUuid is method to generate UUID as string
func GenerateUuid() string {
	return strings.ReplaceAll(uuid.New().String(), "-", "")
}

// CreateBulkResource create new slice from existing resource
func CreateBulkResource(docs []*map[string]interface{}) []*map[string]interface{} {
	var d = make([]*map[string]interface{}, len(docs))
	for i, val := range docs {
		var t = *val
		if _, ok := t["_id"]; !ok {
			t["_id"] = GenerateUuid()
		}
		t["created_at"] = time.Now().UTC().Format(TimeFormat)
		d[i] = &t
	}
	return d
}

// RestGet is method to doing http get request
func RestGet(url string) (int, []byte, error) {
	return fasthttp.GetTimeout(nil, url, TimeOutDuration)
}
