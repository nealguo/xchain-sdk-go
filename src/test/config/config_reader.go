package config

import (
	"io/ioutil"
	"log"
	"xchain-sdk-go/src/test/yaml"
)

var Conf *Yaml

func ReadYaml() {
	conf := new(Yaml)
	yamlFile, err := ioutil.ReadFile("./src/test/res/application.yml")
	if err != nil {
		log.Fatalf("read yaml file wrong, err:%v", err)
	}
	err = yaml.Unmarshal(yamlFile, conf)
	if err != nil {
		log.Fatalf("read config from yaml wrong, err:%v", err)
	}
	log.Println("conf", conf)
	Conf = conf
}
