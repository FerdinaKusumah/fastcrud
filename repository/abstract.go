package repository

import (
	"github.com/FerdinaKusumah/fastcrud/datastore"
	"github.com/spf13/viper"
	"github.com/tidwall/buntdb"
)

type AbstractRepository interface {
	GetList(key string) ([]*map[string]interface{}, error)
	GetDetail(key, id string) ([]*map[string]interface{}, error)
	Set(key string, data []*map[string]interface{}) error
	Delete(key, id string) error
}

type Repository struct {
	AbstractRepository
	MemoryDatastore *buntdb.DB
}

func NewRepository() *Repository {
	var repo = new(Repository)
	if viper.GetString("app.db") == "memory" {
		repo.MemoryDatastore = datastore.NewMemoryDatastore()
	}
	return repo
}
