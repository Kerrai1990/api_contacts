package models

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db *gorm.DB

// GetDB -
func GetDB() *gorm.DB {
	return db
}

func init() {

	e := godotenv.Load()
	if e != nil {
		fmt.Println(e)
	}

	// host := getEnv("DB_HOST", "")
	// username := getEnv("DB_USER", "")
	// password := getEnv("DB_PASS", "")
	// name := getEnv("DB_NAME", "")

	host := "localhost"
	username := "root"
	password := "password"
	name := "contacts"

	connStr := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", host, username, name, password)

	fmt.Println(connStr)

	conn, err := gorm.Open("postgres", connStr)
	if err != nil {
		fmt.Print(err)
	}

	// Migrate DB
	conn.Debug().AutoMigrate(&Account{}, &Contact{})

	defer db.Close()

	fmt.Println("Successfully connected to DB!")

}

// Simple helper function to read an environment or return a default value
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
