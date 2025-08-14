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
	password := GetEnv("DB_PASSWORD", "password")
	host := GetEnv("DB_HOST", "localhost")
	port := GetEnv("DB_PORT", "3360")
	name := GetEnv("DB_NAME", "sample")

	link := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?tls=true", user, password, host, port, name)

	database, err := gorm.Open(mysql.Open(link), &gorm.Config{})
	if err != nil {
		log.Fatalf("Unable to connect with database: %s\n", err)
	} else {
		Database = database
		log.Println("MySQL database is connected successfully ...")
	}
}
