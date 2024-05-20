package config

import (
	"fmt"
	"log"
	"os"
)

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
	TimeZone string
}

type Admin struct {
	Username string
	Password string
}

type AdminInfo struct {
	Admins []Admin
}

type AppConfig struct {
	DatabaseConfig DatabaseConfig
	AdminInfo      AdminInfo
	Secret         []byte
}

var Config *AppConfig

func Init(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Error opening admin credentials file:", err)
	}
	defer file.Close()

	var dbUser, dbPassword, dbName, adminUsername, adminPassword, serviceSecret string
	_, err = fmt.Fscanf(file, "admin_username: %s\nadmin_password: %s\nservice_secret: %s\ndb_user: %s\ndb_password: %s\ndb_name: %s\n",
		&adminUsername, &adminPassword, &serviceSecret, &dbUser, &dbPassword, &dbName)
	if err != nil {
		log.Fatal("Error reading admin credentials:", err)
	}

	Config = &AppConfig{
		Secret: []byte(serviceSecret),
		DatabaseConfig: DatabaseConfig{
			Host:     "localhost",
			Port:     "5432",
			User:     dbUser,
			Password: dbPassword,
			Name:     dbName,
			SSLMode:  "disable",
			TimeZone: "Asia/Novosibirsk",
		},
		AdminInfo: AdminInfo{
			Admins: []Admin{
				{
					Username: adminUsername,
					Password: adminPassword,
				},
			},
		},
	}
}
