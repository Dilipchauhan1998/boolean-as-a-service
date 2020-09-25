package conn

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

//DB store connection to the mysql database
var DB *gorm.DB

//ConnectDB connects to the mysql database, exited if error encountered
func ConnectDB() {
	dbHost := os.Getenv("MYSQL_DB_HOST")
	dbUser := os.Getenv("MYSQL_DB_USER")
	dbPass := os.Getenv("MYSQL_DB_PASS")
	dbName := os.Getenv("MYSQL_DB_NAME")

	dbURL := fmt.Sprintf(
		"%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local",
		dbUser,
		dbPass,
		dbHost,
		dbName,
	)

	db, err := gorm.Open("mysql", dbURL)

	if err != nil {
		fmt.Printf("failed to connect to database!,%v\n", err)
		os.Exit(1)
	}

	DB = db
}
