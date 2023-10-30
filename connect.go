package digital_ocean_db_connect

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type DBConfig struct {
	Username string `json:"username"`
	Password string `json:"password"`
	DBHost   string `json:"dbhost"`
	Port     string `json:"port"`
	DB_Name  string `json:"dbname"`
}

func loadDBConfig(filename string) (*DBConfig, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	config := &DBConfig{}
	err = json.NewDecoder(file).Decode(config)
	if err != nil {
		return nil, err
	}
	return config, nil
}

func buildConnectionString(config *DBConfig) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Username, config.Password, config.DBHost, config.Port, config.DB_Name)
}

func openDatabaseConnection(connectionString string) (*gorm.DB, error) {
	db, err := gorm.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}
	return db, nil
}

//returns db instance
func Connect() *gorm.DB {
	config, err := loadDBConfig("db_info.json")
	if err != nil {
		log.Printf("Error reading DB config: %s\n", err)
		return nil
	}

	connectionString := buildConnectionString(config)
	db, err := openDatabaseConnection(connectionString)
	if err != nil {
		log.Printf("Error Connecting to DB: %s\n", err)
		return nil
	}

	log.Println("Connection to DB established")
	return db
}
