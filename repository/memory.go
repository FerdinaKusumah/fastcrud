package repository

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"github.com/tidwall/buntdb"
)

func (r *Repository) MemoryGetList(key string) ([]*map[string]interface{}, error) {
	var result []*map[string]interface{}
	err := r.MemoryDatastore.View(func(tx *buntdb.Tx) error {
		tx.AscendKeys(fmt.Sprintf(`%s:*`, key), func(key, value string) bool {
			var d = new(map[string]interface{})
			json.Unmarshal([]byte(value), d)
			result = append(result, d)
			return true
		})
		return nil
	})
	return result, err
}

func (r *Repository) MemoryGetDetail(key, id string) (*map[string]interface{}, error) {
	var d = new(map[string]interface{})
	err := r.MemoryDatastore.View(func(tx *buntdb.Tx) error {
		res, err := tx.Get(fmt.Sprintf(`%s:%s`, key, id))
		if err != nil {
			return err
		}
		json.Unmarshal([]byte(res), d)
		return nil
	})
	return d, err
}

func (r *Repository) MemorySet(key string, data []*map[string]interface{}) error {
	// create indexing while inserting data
	var opt = viper.Get("app.opt").(*buntdb.SetOptions)
	r.MemoryDatastore.CreateIndex(key, key+":*", buntdb.IndexString)
	// then update resource
	return r.MemoryDatastore.Update(func(tx *buntdb.Tx) error {
		for _, val := range data {
			t := *val
			res, _ := json.Marshal(val)
			if _, _, err := tx.Set(fmt.Sprintf(`%s:%s`, key, t["_id"]), string(res), opt); err != nil {
				return err
			}
		}
		return nil
	})
}

func (r *Repository) MemoryDelete(key, id string) error {
	err := r.MemoryDatastore.View(func(tx *buntdb.Tx) error {
		if _, err := tx.Delete(fmt.Sprintf(`%s:%s`, key, id)); err != nil {
			return err
		}
		return nil
	})
	return err
}
