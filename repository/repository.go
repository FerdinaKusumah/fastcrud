package repository

import (
	"github.com/spf13/viper"
)

func (r *Repository) GetList(key string) ([]*map[string]interface{}, error) {
	switch viper.GetString("app.db") {
	case "memory":
		return r.MemoryGetList(key)
	default:
		return nil, nil
	}
}

func (r *Repository) GetDetail(key, id string) (*map[string]interface{}, error) {
	switch viper.GetString("app.db") {
	case "memory":
		return r.MemoryGetDetail(key, id)
	default:
		return nil, nil
	}
}

func (r *Repository) Set(key string, data []*map[string]interface{}) error {
	switch viper.GetString("app.db") {
	case "memory":
		return r.MemorySet(key, data)
	default:
		return nil
	}
}

func (r *Repository) Delete(key, id string) error {
	switch viper.GetString("app.db") {
	case "memory":
		return r.MemoryDelete(key, id)
	default:
		return nil
	}
}
