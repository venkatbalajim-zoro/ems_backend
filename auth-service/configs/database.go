package configs

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB

func ConnectDB() {
	user := GetEnv("DB_USER", "root")
	pass := GetEnv("DB_PASSWORD", "password")
	host := GetEnv("DB_HOST", "localhost")
	port := GetEnv("DB_PORT", "3306")
	name := GetEnv("DB_NAME", "sample")

	link := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, pass, host, port, name)

	db, err := gorm.Open(mysql.Open(link), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to MySQL: %s\n", err)
	}

	Database = db
	log.Println("MySQL database is connected successfully ...")
}
