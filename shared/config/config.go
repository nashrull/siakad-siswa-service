package config

import "fmt"

type Config struct {
	DB DBConfig
}

type DBConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Name     string
	Driver   string
}

func (db DBConfig) CreateConnection() string {
	connectionString := ""
	switch db.Driver {
	case "mysql":
		//  "root:@tcp(localhost:3306)/averin?charset=utf8&parseTime=True&loc=Local",
		connectionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", db.Username, db.Password, db.Host, db.Port, db.Name)
	}

	return connectionString
}
