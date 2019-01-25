package config

import (
	"database/sql"
	"github.com/mongodb/mongo-go-driver/mongo"
	"net/http"
)

const (
	MySecret = "SenelisMegstaMociutesApvalumus"
)

var (
	Flags FlagSettings
	Conf  *Config

	db         *sql.DB
	dbUri      string
	driverName string

	mongoDB *mongo.Database
)

type (
	DbConfig struct {
		Server     string `yaml:"server"`
		DriverName string `yaml:"driverName"`
		DbName     string `yaml:"dbname,omitempty"`
		User       string `yaml:"user,omitempty"`
		Password   string `yaml:"password,omitempty"`
		Charset    string `yaml:"charset,omitempty"`
	}
	MongoDBConfig struct {
		Connection string `yaml:"connection"`
		DbName     string `yaml:"db_name"`
	}

	Config struct {
		Port       string        `yaml:"port"`
		AssetsPath string        `yaml:"assets"`
		LogFile    string        `yaml:"logFile"`
		Host       string        `yaml:"host"`
		Db         DbConfig      `yaml:"db"`
		MongoDB    MongoDBConfig `yaml:"mongoDB"`
	}

	FlagSettings struct {
		Config     string
		AssetsPath string
		Assets     http.Dir
		Host       string
		Port       string
		LogFile    string
	}
)
