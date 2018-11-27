package core

import (
	fr "github.com/DATA-DOG/fastroute"
)

type ApiModule struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Version     string `yaml:"version"`
	Author      string `yaml:"author"`
}

var Routes map[string]fr.Router

func init() {
	Routes = make(map[string]fr.Router)
}
