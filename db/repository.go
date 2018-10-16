package db

import (
	"../config"
	"../helpers"
)

const sqlShowTables = "SHOW TABLES"
const sqlShowDatabases = "SHOW DATABASES"
const sqlGetDbServerVersion = "SELECT version()"

func GetVersion() (string, error) {
	db, err := config.Conf.GetDb()
	helpers.CheckErr(err)

	row := db.QueryRow(sqlGetDbServerVersion)
	var version string
	err = row.Scan(&version)
	helpers.CheckErr(err)
	return version, nil
}

func showDatabases() ([]string, error) {
	db, err := config.Conf.GetDb()
	helpers.CheckErr(err)

	databases := make([]string, 0)

	res, _ := db.Query(sqlShowDatabases)

	for res.Next() {
		var database string
		err = res.Scan(&database)
		helpers.CheckErr(err)
		databases = append(databases, database)
	}
	return databases, nil
}

func showTables() ([]string, error) {

	db, err := config.Conf.GetDb()
	helpers.CheckErr(err)

	tables := make([]string, 0)

	res, _ := db.Query(sqlShowTables)

	for res.Next() {
		var table string
		err = res.Scan(&table)
		helpers.CheckErr(err)
		tables = append(tables, table)
	}
	return tables, nil
}
