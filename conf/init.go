package conf

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

const projectDir = "/Users/bytedance/go/src/github.com/JackTJC/backend/"

var Conf *YamlConf
type YamlConf struct {
	Server struct{
		Port string `yaml:"port"`
	} `yaml:"server"`
	Mysql struct{
		UserName string `yaml:"user_name"`
		Passwd string `yaml:"passwd"`
		Host string`yaml:"host"`
		Port string `yaml:"port"`
		Database string `yaml:"database"`
	} `yaml:"mysql"`
	Redis struct{
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"redis"`
}

func InitConf(){
	Conf=&YamlConf{}
	file,err:=ioutil.ReadFile(projectDir+"conf/local.yaml")
	if err!=nil{
		panic(err)
	}
	err=yaml.Unmarshal(file,Conf)
	if err!=nil{
		panic(err)
	}
}
