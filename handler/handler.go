package handler

import (
	"encoding/json"
	"github.com/FerdinaKusumah/fastcrud/model"
	"github.com/FerdinaKusumah/fastcrud/utils"
	"github.com/gorilla/mux"
	"net/http"
)

// HealthCheck method to determining api is live
func HealthCheck(out http.ResponseWriter, in *http.Request) {
	_ = json.NewEncoder(out).Encode(utils.HttpMessage["live"])
}

// LoadExampleApi method to determining api is live
func LoadExampleApi(out http.ResponseWriter, in *http.Request) {
	var (
		result = new(model.Response)
		err    error
	)
	if result, err = GetUseCase().LoadExampleData(); err != nil {
		_ = json.NewEncoder(out).Encode(result)
		return
	}
	_ = json.NewEncoder(out).Encode(result)
}

// ListResource method to determining api is live
func ListResource(out http.ResponseWriter, in *http.Request) {
	var (
		result = new(model.Response)
		param  = mux.Vars(in)
		key    = param["key"]
		err    error
	)
	if result, err = GetUseCase().ListResource(key); err != nil {
		_ = json.NewEncoder(out).Encode(result)
		return
	}
	_ = json.NewEncoder(out).Encode(result)
}

// CreateResource method to determining api is live
func CreateResource(out http.ResponseWriter, in *http.Request) {
	var (
		payload = make(map[string]interface{})
		result  = new(model.Response)
		param   = mux.Vars(in)
		key     = param["key"]
		err     error
	)
	if err = json.NewDecoder(in.Body).Decode(&payload); err != nil {
		_ = json.NewEncoder(out).Encode(utils.HttpMessage["invalid_json"])
		return
	}
	if result, err = GetUseCase().CreateResource(key, payload); err != nil {
		_ = json.NewEncoder(out).Encode(utils.HttpMessage["failed_created"])
		return
	}
	_ = json.NewEncoder(out).Encode(result)
}

// CreateBulkResource method to determining api is live
func CreateBulkResource(out http.ResponseWriter, in *http.Request) {
	var (
		result  = new(model.Response)
		payload []*map[string]interface{}
		param   = mux.Vars(in)
		key     = param["key"]
		err     error
	)
	if err = json.NewDecoder(in.Body).Decode(&payload); err != nil {
		_ = json.NewEncoder(out).Encode(utils.HttpMessage["invalid_json"])
		return
	}
	if result, err = GetUseCase().CreateBulkResource(key, payload); err != nil {
		_ = json.NewEncoder(out).Encode(utils.HttpMessage["failed_created"])
		return
	}
	_ = json.NewEncoder(out).Encode(result)
}

// DetailResource method to determining api is live
func DetailResource(out http.ResponseWriter, in *http.Request) {
	var (
		result = new(model.Response)
		param  = mux.Vars(in)
		key    = param["key"]
		id     = param["id"]
		err    error
	)
	if result, err = GetUseCase().DetailResource(key, id); err != nil {
		_ = json.NewEncoder(out).Encode(result)
		return
	}
	_ = json.NewEncoder(out).Encode(result)
}

// UpdateResource method to determining api is live
func UpdateResource(out http.ResponseWriter, in *http.Request) {
	var (
		payload = make(map[string]interface{})
		result  = new(model.Response)
		param   = mux.Vars(in)
		key     = param["key"]
		id      = param["id"]
		err     error
	)
	if err = json.NewDecoder(in.Body).Decode(&payload); err != nil {
		_ = json.NewEncoder(out).Encode(utils.HttpMessage["invalid_json"])
		return
	}
	if result, err = GetUseCase().UpdateResource(key, id, payload); err != nil {
		_ = json.NewEncoder(out).Encode(utils.HttpMessage["failed_updated"])
		return
	}
	_ = json.NewEncoder(out).Encode(result)
}

// DeleteResource method to determining api is live
func DeleteResource(out http.ResponseWriter, in *http.Request) {
	var (
		result = new(model.Response)
		param  = mux.Vars(in)
		key    = param["key"]
		id     = param["id"]
		err    error
	)
	if result, err = GetUseCase().DeleteResource(key, id); err != nil {
		_ = json.NewEncoder(out).Encode(utils.HttpMessage["failed_deleted"])
		return
	}
	_ = json.NewEncoder(out).Encode(result)
}
