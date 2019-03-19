package gocfg

import (
	"errors"
	"gopkg.in/ini.v1"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

func Load(filePath string, gcfg interface{}) error {
	var (
		conf *ini.File
		err  error
	)
	if _, err := os.Stat(filePath); err != nil && os.IsExist(err) {
		return errors.New("file not found")
	}
	if conf, err = ini.Load(filePath); err != nil {
		return errors.New("Loading config ini error : " + err.Error())
	}

	if err = conf.MapTo(gcfg); err != nil {
		return errors.New("map to struct error : " + err.Error())
	}
	return nil
}

func LoadYml(filepath string, gcfg interface{}) error {
	var (
		text []byte
		err  error
	)
	text, err = ioutil.ReadFile(filepath)
	if err != nil {
		return err
	}
	if err = yaml.Unmarshal(text, gcfg); err != nil {
		return err
	}
	return nil
}
