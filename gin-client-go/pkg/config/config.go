package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
	"k8s.io/klog"
)

type Config struct {
	Server Server
}

type Server struct {
	Name string `yaml:"name"`
	Host string `yanl:"host"`
	Port string `yaml:"port"`
}

func init() {
	var config Config
	yamlFile, err := ioutil.ReadFile(".config.yaml")
	if err != nil {
		klog.Fatal("read config error", err)
		return
	}

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		klog.Fatal("unmashal error", err)
		return
	}
}


func 