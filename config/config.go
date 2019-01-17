package config

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/mongodb/mongo-go-driver/mongo"
	"go-api-ws/helpers"
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
type MongoDBConfig struct {
	Connection string `yaml:"connection"`
	DbName     string `yaml:"db_name"`
}

type Config struct {
	Port       string        `yaml:"port"`
	AssetsPath string        `yaml:"assets"`
	LogFile    string        `yaml:"logFile"`
	Host       string        `yaml:"host"`
	Db         DbConfig      `yaml:"db"`
	MongoDB    MongoDBConfig `yaml:"mongoDB"`
}

type FlagSettings struct {
	Config     string
	AssetsPath string
	Assets     http.Dir
	Host       string
	Port       string
	LogFile    string
}

const MySecret = "SenelisMegstaMociutesApvalumus"

var Flags FlagSettings
var Conf *Config

var db *sql.DB
var dbUri string
var driverName string

var mongoDB *mongo.Database

func (c *Config) GetMongoDb() *mongo.Database {
	client, err := mongo.NewClient(c.MongoDB.Connection)
	helpers.PanicErr(err)

	err = client.Connect(context.Background())
	helpers.PanicErr(err)

	// Collection types can be used to access the database
	mongoDB = client.Database(c.MongoDB.DbName)

	return mongoDB
}

func NewFlags() *FlagSettings {
	Flags := FlagSettings{}
	return &Flags
}

func GetConfig(configFile string) *Config {
	Conf = &Config{}
	if configFile != "" {
		err := Conf.GetConfFromFile(configFile)
		helpers.PanicErr(err)
	}
	return Conf
}

func (c *Config) GetDb() (*sql.DB, error) {
	if db == nil {
		if dbUri == "" {
			dbConfig := c.Db
			driverName = dbConfig.DriverName
			dbUri = fmt.Sprintf("%s:%s@%s/%s?charset=%s&parseTime=True",
				dbConfig.User, dbConfig.Password,
				dbConfig.Server, dbConfig.DbName, dbConfig.Charset)
		}
		ldb, err := sql.Open(driverName, dbUri)
		if err != nil {
			return nil, err
		}
		db = ldb
	}
	return db, nil
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
