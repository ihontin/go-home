package main

import (
	"gopkg.in/yaml.v3"
)

// getConfigFromYAML, которая будет получать данные из YAML и возвращать структуру Config и ошибку, если возникнет.
func getConfigFromYAML(data []byte) (Config, error) {
	var deserConf Config
	err := yaml.Unmarshal(data, &deserConf)
	if err != nil {
		return Config{}, err
	}
	return deserConf, nil
}

type Config struct {
	Server Server `yaml:"server"`
	Db     Db     `yaml:"db"`
}

type Server struct {
	Port string `yaml:"port"`
}

type Db struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

//func main() {
//	var servData = []byte(`server:
//  port: 8080
//db:
//  host: localhost
//  port: 5432
//  user: admin
//  password: password123
//`)
//	typePrint, err := getConfigFromYAML(servData)
//	if err != nil {
//		fmt.Errorf("error: %v", err)
//	}
//	fmt.Println(typePrint.Db.)
//}

//- server:
//port: "8080"
//db:
//host: localhost
//port: "5432"
//user: admin
//password: password123
