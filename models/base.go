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

func init() {

	connStr := "root:password@tcp(127.0.0.1:33060)/contacts?charset=utf8"

	var err error

	db, err = gorm.Open("mysql", connStr)
	if err != nil {
		fmt.Println(connStr)
		fmt.Println(err)
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
