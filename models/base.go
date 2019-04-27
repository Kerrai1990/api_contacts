package models

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

// GetDB -
func GetDB() *gorm.DB {
	return db
}

// MakeDB -
func MakeDB() {

	fmt.Println("Connecting to DB")
	connStr := "root:password@tcp(api_contacts_db:3306)/contacts?charset=utf8"
	fmt.Println(connStr)

	var err error

	db, err = gorm.Open("mysql", connStr)
	if err != nil {
		fmt.Println(connStr)
		fmt.Println(err)
		return
	}

	fmt.Println("Successfully connected to DB!")

	// Migrate DB
	fmt.Println("Running Migration")
	db.Debug().AutoMigrate(&Account{}, &Contact{})
	fmt.Println("Migration Complete")
}

// Simple helper function to read an environment or return a default value
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
