package mongodb

import (
	"SoftWeb/lib/jlog"
	"SoftWeb/lib/pool"
	"errors"
	"fmt"
	"time"

	"gopkg.in/mgo.v2/bson"
)

type DataConfig struct {
	Id         string            `bson:"_id"`
	Channel    string            `bson:"channel"`
	DataType   string            `bson:"dateType"`
	UpdateTime time.Time         `bson:"updateTime"`
	Data       map[string]string `bson:"data"`
}

const (
	DATA_CONFIG_COLLECTION string = "dataConfig"
)

func GetDataConfig(channel, dataType string,
	keys []string) (map[string]string, error) {

	defer func() {
		if re := recover(); re != nil {
			err := fmt.Errorf("%v", re)
			jlog.LogErrorSendMail(err)
		}
	}()
	var data DataConfig
	result := make(map[string]string)
	err := MgoPool.Call(func(src pool.Src) error {
		c := src.(*MgoSrc).DB(mongoConf.DB).C(DATA_CONFIG_COLLECTION)
		err := c.Find(bson.M{}).One(data)
		if err != nil {
			return err
		}
		for _, key := range keys {
			value, ok := data.Data[key]
			if ok {
				result[key] = value
			}
		}
		return nil
	})
	return result, err
}

func SetDataConfig(channel, dataType string, data map[string]string) error {
	defer func() {
		if re := recover(); re != nil {
			err := fmt.Errorf("%v", re)
			jlog.LogErrorSendMail(err)
		}
	}()
	err := MgoPool.Call(func(src pool.Src) error {
		//c := src.(*MgoSrc).DB(mongoConf.DB).C(DATA_CONFIG_COLLECTION)
		return errors.New("ss")
	})
	return err
}
