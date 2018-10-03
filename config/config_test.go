package config

import (
	"fmt"
	"testing"
)

var s = `
port: :8080
db:
  driverName: mysql
  user: new_user
  password: secretPass
  dbname:   express_special
  charset:  utf8
  server:   "tcp(255.155.123.222:3306)"`

func TestReadYamlString(t *testing.T) {
	testing.Short()
	c, err := ReadYamlString(s)
	if err != nil {
		fmt.Printf("Err: %s", err)
	}
	fmt.Printf("c %s", c)

	if c == nil {
		t.Errorf("returned nil")
	}

	expectedPort := ":8080"
	if c.Port != expectedPort {
		t.Errorf("Port value retrieved: %s expeted %s", c.Port, expectedPort)
	}

	expectedDriverName := "mysql"
	if c.Db.DriverName != expectedDriverName {
		t.Errorf("expected \"db.driverName\" value: %s, actual value: %s", expectedDriverName, c.Db.DriverName)
	}

	expectedUserNameName := "new_user"
	if c.Db.User != expectedUserNameName {
		t.Errorf("expected \"db.username\" value: %s, actual value: %s", expectedUserNameName, c.Db.User)
	}

	expectedPassword := "secretPass"
	if c.Db.Password != expectedPassword {
		t.Errorf("expected \"db.password\" value: %s, actual value: %s", expectedPassword, c.Db.Password)
	}

	expectedDbName := "express_special"
	if c.Db.DbName != expectedDbName {
		t.Errorf("expected \"db.dbName\" value: %s, actual value: %s", expectedDbName, c.Db.DbName)
	}

	expectedCharset := "utf8"
	if c.Db.Charset != expectedCharset {
		t.Errorf("expected \"db.charset\" value: '%s', actual value: '%s'", expectedCharset, c.Db.Charset)
	}

	expectedServer := "tcp(255.155.123.222:3306)"
	if c.Db.Server != expectedServer {
		t.Errorf("expected \"db.server\" value: %s, actual value: %s", expectedServer, c.Db.Server)
	}
}
