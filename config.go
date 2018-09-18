package gocfg

import (
	`errors`
	`gopkg.in/ini.v1`
	`os`
)

func Load(filePath string, gcfg interface{}) error {
	var conf *ini.File
	if _, err := os.Stat(filePath); err != nil && os.IsExist(err) {
		return errors.New("file not found")
	}
	conf, err := ini.Load(filePath)
	if err != nil {
		return errors.New("Loading config ini error : " + err.Error())
	}
	err = conf.MapTo(gcfg)
	if err != nil {
		return errors.New("map to struct error : " + err.Error())
	}
	return nil
}
