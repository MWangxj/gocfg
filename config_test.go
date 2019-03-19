package gocfg

import (
	`github.com/MWangxj/logger`
	`reflect`
	`testing`
)

func printInter(gcfg interface{}) {
	t := reflect.ValueOf(gcfg).Elem()
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		logger.Info(t.Type().Field(i).Name, f.Interface())
	}
}

type Config struct {
	Port int `ini:"port"`
}

func TestLoadConfig(t *testing.T) {
	c := &Config{}
	Load("./test.yaml", c)
	printInter(c)
}
