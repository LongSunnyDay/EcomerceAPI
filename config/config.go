package config

import (
	"database/sql"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type DbConfig struct {
	Server     string `yaml:"server"`
	DriverName string `yaml:"driverName"`
	DbName     string `yaml:"dbname,omitempty"`
	User       string `yaml:"user,omitempty"`
	Password   string `yaml:"password,omitempty"`
	Charset    string `yaml:"charset,omitempty"`
}

type Config struct {
	Port string `yaml:"port"`
	AssetsPath string `yaml:"assets"`
	LogFile string `yaml:"logFile"`
	Host string `yaml:"host"`
	Db   DbConfig `yaml:"db"`
}

type FlagSettings struct {
	Config string
	AssetsPath string
	Assets http.Dir
	Host string
	Port string
	LogFile string
}

var Flags FlagSettings
var Conf *Config

var dbUri string
var driverName string

func NewFlags() *FlagSettings {
	Flags := FlagSettings{}
	return &Flags
}

func GetConfig(configFile string) *Config {
	Conf = &Config{}
	if configFile != "" {
		Conf.GetConfFromFile(configFile)
	}
	return Conf
}

func (c *Config) GetDb() (*sql.DB, error) {
	if dbUri == "" {
		dbConfig := c.Db
		driverName = dbConfig.DriverName
		dbUri = fmt.Sprintf("%s:%s@%s/%s?charset=%s&parseTime=True",
			dbConfig.User, dbConfig.Password,
			dbConfig.Server, dbConfig.DbName, dbConfig.Charset)

	}
	return sql.Open(driverName, dbUri)
}

func (c *Config) GetConfFromFile(fileName string) error {
	pwd, _ := os.Getwd()
	yamlFile, err := ioutil.ReadFile(pwd + "/" + fileName)
	if err != nil {
		log.Printf("%s file read error.  #%v\n", fileName, err)
	}
	return c.GetConfFromString(string(yamlFile))
}

func (c *Config) GetConfFromString(yamlString string) error {

	err := yaml.Unmarshal([]byte(yamlString), c)
	if err != nil {
		log.Fatalf("%s parse error %v\n", yamlString, err)
	}
	return err
}
