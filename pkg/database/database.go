package database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jonatascabral/zipcodes-api/pkg/models"
	"log"
)

type Config struct {
	Host string
	Port int
	DbName string
	User string
	Password string
	Charset string
	DB *gorm.DB
}

var database *gorm.DB

func Connect(driver string, config *Config) *gorm.DB {
	var connectionString string

	switch driver {
	case "mysql":
		connectionString = fmt.Sprintf(
			"%s:%s@(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
			config.User,
			config.Password,
			config.Host,
			config.Port,
			config.DbName,
			config.Charset)
	case "postgres":
		connectionString = fmt.Sprintf(
			"host=%s port=%d user=%s dbname=%s password=%s",
			config.Host,
			config.Port,
			config.User,
			config.DbName,
			config.Password)
	}

	var err error
	database, err = gorm.Open(driver, connectionString)
	if err != nil {
		panic(err)
	}
	database.LogMode(true)
	return database
}

func Close() {
	database.Close()
}

func Migrate(models *models.Address) {
	database.AutoMigrate(&models)
}

func Init() *Config {
	config := &Config{
		User: "root",
		Port: 3306,
		Password: "root",
		Host: "localhost",
		DbName: "zipcodes_api",
		Charset: "utf8"}
	config.DB = Connect("mysql", config)
	log.Println("Connected to database")

	log.Println("Migrating database")
	Migrate(&models.Address{})
	log.Println("Database migrated")

	return config
}
