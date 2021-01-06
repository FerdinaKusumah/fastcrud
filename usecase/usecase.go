package usecase

import (
	"encoding/json"
	"github.com/FerdinaKusumah/fastcrud/model"
	"github.com/FerdinaKusumah/fastcrud/repository"
	"github.com/FerdinaKusumah/fastcrud/utils"
	"net/http"
	"time"
)

type UseCase struct {
	Repository *repository.Repository
}

func NewUseCase(repo *repository.Repository) *UseCase {
	return &UseCase{
		Repository: repo,
	}
}

// LoadExampleData ...
func (uc *UseCase) LoadExampleData() (*model.Response, error) {
	var (
		res = new(model.Response)
		ch  = make(chan struct{}, len(utils.ExampleJsonData))
	)
	defer close(ch)

	for key, val := range utils.ExampleJsonData {
		// use go routine
		go func(ch chan<- struct{}, key, url string) {
			var t []*map[string]interface{}
			_, data, _ := utils.RestGet(url)
			if err := json.Unmarshal(data, &t); err != nil {
				utils.SendLogError("unable to parse json", err)
			}
			if err := uc.Repository.Set(key, utils.CreateBulkResource(t)); err != nil {
				utils.SendLogError("unable to insert data example", err)
			}
			ch <- struct{}{}
		}(ch, key, val)
	}

	for i := 0; i < len(utils.ExampleJsonData); i++ {
		<-ch
	}

	res.Status = http.StatusOK
	res.Message = utils.StatusMessage["successfully_seed_mock"]
	return res, nil
}

// ListResource ...
func (uc *UseCase) ListResource(key string) (*model.Response, error) {
	var res = new(model.Response)
	res.Data, _ = uc.Repository.GetList(key)
	res.Status = http.StatusOK
	return res, nil
}

// CreateResource ...
func (uc *UseCase) CreateResource(key string, data map[string]interface{}) (*model.Response, error) {
	var res = new(model.Response)
	if _, ok := data["_id"]; !ok {
		data["_id"] = utils.GenerateUuid()
	}
	data["created_at"] = time.Now().UTC().Format(utils.TimeFormat)
	if err := uc.Repository.Set(key, []*map[string]interface{}{&data}); err != nil {
		return nil, err
	}
	res.Data = data
	res.Status = http.StatusCreated
	return res, nil
}

// CreateBulkResource ...
func (uc *UseCase) CreateBulkResource(key string, data []*map[string]interface{}) (*model.Response, error) {
	var res = new(model.Response)
	if err := uc.Repository.Set(key, utils.CreateBulkResource(data)); err != nil {
		return nil, err
	}
	res.Status = http.StatusCreated
	res.Message = utils.StatusMessage["successfully_created"]
	return res, nil
}

// DetailResource ...
func (uc *UseCase) DetailResource(key, id string) (*model.Response, error) {
	var res = new(model.Response)
	res.Data, _ = uc.Repository.GetDetail(key, id)

	res.Status = http.StatusOK
	return res, nil
}

// UpdateResource ...
func (uc *UseCase) UpdateResource(key, id string, data map[string]interface{}) (*model.Response, error) {
	var res = new(model.Response)
	data["_id"] = id

	if err := uc.Repository.Set(key, []*map[string]interface{}{&data}); err != nil {
		return nil, err
	}

	res.Status = http.StatusCreated
	res.Data = data
	return res, nil
}

// DeleteResource ...
func (uc *UseCase) DeleteResource(key, id string) (*model.Response, error) {
	var res = new(model.Response)
	if err := uc.Repository.Delete(key, id); err != nil {
		return nil, err
	}

	res.Status = http.StatusNoContent
	return res, nil
}
