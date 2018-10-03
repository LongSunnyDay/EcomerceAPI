package user

import "go-api-ws/core"

var userModule core.ApiModule

func init() {
	userModule = core.ApiModule{
		Name: "User module",
		Description: "User module. Supports username and email authentication. Categories are stored as a flat list.",
		Version: "0.1",
		Author: "Remigijus Bauzys @ JivaLabs",
	}
}