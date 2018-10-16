package auth

import (
	"../config"
	"../helpers"
	_ "database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var ErrUserNotFound = errors.New("User not found!")

func CreateUsersTableIfNotExists() {
	db, err := config.Conf.GetDb()
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(`
    CREATE TABLE IF NOT EXISTS jiva_user (
      id INT UNSIGNED NOT NULL AUTO_INCREMENT,
      username VARCHAR(32) NULL,
      password VARCHAR(512) NULL,
    PRIMARY KEY (id)
    )`)
	helpers.CheckErr(err)
}

